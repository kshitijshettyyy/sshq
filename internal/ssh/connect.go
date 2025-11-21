package ssh

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

func ConnectToServer(username, address, password, pemFile string) error {
	var auth ssh.AuthMethod

	if pemFile != "" {
		key, err := ioutil.ReadFile(pemFile)
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return err
		}
		auth = ssh.PublicKeys(signer)
	} else {
		auth = ssh.Password(password)
	}

	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{auth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", address+":22", config)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	// Set up terminal
	fd := int(os.Stdin.Fd())
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return err
	}
	defer func() {
		term.Restore(fd, oldState)
		fmt.Print("\r") // reset cursor to beginning of line
	}()

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", 80, 40, modes); err != nil {
		return fmt.Errorf("request for pseudo terminal failed: %v", err)
	}

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	// Handle Ctrl+C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		for range sigs {
			session.Signal(ssh.SIGINT)
		}
	}()

	// fmt.Printf("Last login: connected successfully.\n")
	if err := session.Shell(); err != nil {
		return err
	}

	err = session.Wait()
	if err != nil {
		fmt.Printf("\rDisconnected.\n")
	} else {
		fmt.Printf("\rDisconnected gracefully.\n")
	}
	return nil
}
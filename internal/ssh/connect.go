package ssh

import (
	"fmt"
	"io/ioutil"
	// "log"
	"os"

	"golang.org/x/crypto/ssh"
	"sshq/internal/vault"
)

func Connect(entry *vault.Entry) error {
	var auth ssh.AuthMethod

	if entry.Method == "password" {
		auth = ssh.Password(entry.Password)
	} else if entry.Method == "pem" {
		key, err := ioutil.ReadFile(entry.PEMFile)
		if err != nil {
			return err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return err
		}
		auth = ssh.PublicKeys(signer)
	} else {
		return fmt.Errorf("unknown auth method: %s", entry.Method)
	}

	config := &ssh.ClientConfig{
		User:            entry.User,
		Auth:            []ssh.AuthMethod{auth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", entry.Host+":22", config)
	if err != nil {
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	fmt.Println("✅ Connected! Type 'exit' to disconnect.")
	return session.Shell()
}

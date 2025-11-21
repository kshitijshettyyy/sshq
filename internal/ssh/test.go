package ssh

import (
	"fmt"
	"io/ioutil"
	"time"

	"golang.org/x/crypto/ssh"
	"sshq/internal/vault"
)

// TestConnection tries connecting and running a simple command (like `echo ok`)
// to verify that credentials are valid.
func TestConnection(entry *vault.Entry) error {
	var auth ssh.AuthMethod

	if entry.Method == "password" {
		auth = ssh.Password(entry.Password)
	} else if entry.Method == "pem" {
		key, err := ioutil.ReadFile(entry.PEMFile)
		if err != nil {
			return fmt.Errorf("failed to read PEM file: %w", err)
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return fmt.Errorf("invalid PEM key: %w", err)
		}
		auth = ssh.PublicKeys(signer)
	} else {
		return fmt.Errorf("unknown auth method: %s", entry.Method)
	}

	config := &ssh.ClientConfig{
		User:            entry.User,
		Auth:            []ssh.AuthMethod{auth},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", entry.Host+":22", config)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	defer session.Close()

	if err := session.Run("echo ok"); err != nil {
		return fmt.Errorf("command test failed: %w", err)
	}

	return nil
}

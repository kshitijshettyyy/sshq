package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"sshq/config"
	"strings"
)

func ConnectToHost(name string) {
	hosts := config.LoadHosts()

	host, exists := hosts[name]
	if !exists {
		fmt.Printf("❌ Host '%s' not found.\n", name)
		return
	}

	addr := fmt.Sprintf("%s@%s", host.User, host.Address)
	sshCmd := exec.Command("ssh", "-p", fmt.Sprint(host.Port), addr)
	sshCmd.Stdin = os.Stdin
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr

	err := sshCmd.Run()
	if err != nil {
		fmt.Println("⚠️ SSH connection failed. Attempting to set up passwordless access...")

		fmt.Printf("🔐 Run ssh-copy-id to install SSH key on '%s'? (y/n): ", addr)
		var response string
		fmt.Scanln(&response)
		if strings.ToLower(response) != "y" {
			fmt.Println("❌ Aborting setup.")
			return
		}

		copyCmd := exec.Command("ssh-copy-id", "-p", fmt.Sprint(host.Port), addr)
		copyCmd.Stdin = os.Stdin
		copyCmd.Stdout = os.Stdout
		copyCmd.Stderr = os.Stderr

		if err := copyCmd.Run(); err != nil {
			fmt.Println("❌ ssh-copy-id failed. Please check your SSH key and VM credentials.")
			return
		}

		fmt.Println("✅ SSH key copied successfully. Retrying connection...")

		retryCmd := exec.Command("ssh", "-p", fmt.Sprint(host.Port), addr)
		retryCmd.Stdin = os.Stdin
		retryCmd.Stdout = os.Stdout
		retryCmd.Stderr = os.Stderr
		retryCmd.Run()
	}
}

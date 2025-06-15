package main

import (
	"fmt"
	"os"
	"os/exec"
	"sshq/config"
	"sshq/models"
	"sshq/ssh"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		handleAdd()
	case "connect":
		handleConnect()
	case "delete":
		handleDelete()
	case "list":
		handleList()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		printUsage()
	}
}

func handleAdd() {
	if len(os.Args) != 6 || os.Args[4] != "-p" {
		fmt.Println("Invalid format. Use username@address")
		return
	}
	name := os.Args[2]
	userHost := os.Args[3]
	port, err := strconv.Atoi(os.Args[5])
	if err != nil {
		fmt.Println("Invalid port.")
		return
	}
	userHostParts := parseUserHost(userHost)
	if userHostParts == nil {
		fmt.Println("Invalid format. Use username@address")
		return
	}

	host := models.Host{
		User:    userHostParts[0],
		Address: userHostParts[1],
		Port:    port,
	}

	config.SaveHost(name, host)
	fmt.Printf("✅ Host '%s' added.\n", name)

	fmt.Printf("🔐 Run ssh-copy-id to install SSH key on '%s@%s'? (y/n): ", host.User, host.Address)
	var response string
	fmt.Scanln(&response)
	if strings.ToLower(response) == "y" {
		copyCmd := exec.Command("ssh-copy-id", "-p", fmt.Sprint(port), fmt.Sprintf("%s@%s", host.User, host.Address))
		copyCmd.Stdin = os.Stdin
		copyCmd.Stdout = os.Stdout
		copyCmd.Stderr = os.Stderr
		if err := copyCmd.Run(); err != nil {
			fmt.Println("❌ ssh-copy-id failed. You can do it manually with:")
			fmt.Printf("   ssh-copy-id -p %d %s@%s\n", port, host.User, host.Address)
		} else {
			fmt.Println("✅ SSH key copied successfully.")
		}
	}
}

func handleConnect() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: sshq connect <name>")
		return
	}
	ssh.ConnectToHost(os.Args[2])
}

func handleDelete() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: sshq delete <name>")
		return
	}
	name := os.Args[2]
	config.DeleteHost(name)
}

func handleList() {
	hosts := config.LoadHosts()
	if len(hosts) == 0 {
		fmt.Println("📭 No hosts found.")
		return
	}
	fmt.Println("📋 Saved Hosts:")
	for name, h := range hosts {
		fmt.Printf("• %s: %s@%s -p %d\n", name, h.User, h.Address, h.Port)
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  sshq add <name> <user@host> -p <port>")
	fmt.Println("  sshq connect <name>")
	fmt.Println("  sshq delete <name>")
	fmt.Println("  sshq list")
}

func parseUserHost(userHost string) []string {
	parts := make([]string, 2)
	for i, v := range userHost {
		if v == '@' {
			parts[0] = userHost[:i]
			parts[1] = userHost[i+1:]
			return parts
		}
	}
	return nil
}

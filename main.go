package main

import (
	"bufio"
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

	command := os.Args[1]

	switch command {
	case "add":
		addHost()
	case "connect":
		if len(os.Args) != 3 {
			fmt.Println("Usage: sshq connect <name>")
			return
		}
		ssh.ConnectToHost(os.Args[2])
	case "list":
		listHosts()
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: sshq delete <name>")
			return
		}
		config.DeleteHost(os.Args[2])
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  sshq add [<alias> <user@host[:port]>]")
	fmt.Println("  sshq connect <alias>")
	fmt.Println("  sshq list")
	fmt.Println("  sshq delete <alias>")
}

func addHost() {
	var name, user, address string
	port := 22 // default

	if len(os.Args) == 4 {
		name = os.Args[2]
		userHost := os.Args[3]
		parsedUser, parsedHost, parsedPort := parseUserHost(userHost)
		if parsedUser != "" {
			user = parsedUser
		} else {
			user = "root"
		}
		address = parsedHost
		if parsedPort != 0 {
			port = parsedPort
		}
	} else {
		// Interactive mode
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Alias name: ")
		scanner.Scan()
		name = scanner.Text()

		fmt.Print("Host (e.g., 192.168.0.5): ")
		scanner.Scan()
		address = scanner.Text()

		fmt.Print("Username [root]: ")
		scanner.Scan()
		user = scanner.Text()
		if user == "" {
			user = "root"
		}

		fmt.Print("Port [22]: ")
		scanner.Scan()
		portInput := scanner.Text()
		if portInput != "" {
			parsed, err := strconv.Atoi(portInput)
			if err == nil {
				port = parsed
			}
		}
	}

	host := models.Host{
		User:    user,
		Address: address,
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

			// Try to silently add key to ssh-agent
			home, _ := os.UserHomeDir()
			keyPath := fmt.Sprintf("%s/.ssh/id_rsa", home)
			addCmd := exec.Command("ssh-add", keyPath)
			addCmd.Stdin = os.Stdin
			addCmd.Stdout = os.Stdout
			addCmd.Stderr = os.Stderr
			_ = addCmd.Run()
		}
	}
}

func parseUserHost(input string) (user, host string, port int) {
	port = 0
	if strings.Contains(input, "@") {
		parts := strings.SplitN(input, "@", 2)
		user = parts[0]
		input = parts[1]
	} else {
		user = "root"
	}

	if strings.Contains(input, ":") {
		hostParts := strings.SplitN(input, ":", 2)
		host = hostParts[0]
		portParsed, err := strconv.Atoi(hostParts[1])
		if err == nil {
			port = portParsed
		}
	} else {
		host = input
	}
	return
}

func listHosts() {
	hosts := config.LoadHosts()
	if len(hosts) == 0 {
		fmt.Println("📭 No hosts found.")
		return
	}
	fmt.Println("📋 Saved hosts:")
	for name, host := range hosts {
		fmt.Printf("- %s → %s@%s\n", name, host.User, host.Address)
	}
}

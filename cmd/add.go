package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"sshq/internal/ssh"
	"sshq/internal/vault"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new server to your SSH vault",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter alias name (e.g. db-prod): ")
		alias, _ := reader.ReadString('\n')
		alias = strings.TrimSpace(alias)

		fmt.Print("Enter hostname or IP: ")
		host, _ := reader.ReadString('\n')
		host = strings.TrimSpace(host)

		fmt.Print("Enter username: ")
		user, _ := reader.ReadString('\n')
		user = strings.TrimSpace(user)

		fmt.Print("Auth method (password/pem): ")
		method, _ := reader.ReadString('\n')
		method = strings.TrimSpace(method)

		var password, pemPath string
		switch method {
		case "password":
			fmt.Print("Enter password: ")
			password, _ = reader.ReadString('\n')
			password = strings.TrimSpace(password)
		case "pem":
			fmt.Print("Enter PEM file path: ")
			pemPath, _ = reader.ReadString('\n')
			pemPath = strings.TrimSpace(pemPath)
		default:
			fmt.Println("❌ Unsupported auth method.")
			return
		}

		entry := vault.Entry{
			Alias:    alias,
			Host:     host,
			User:     user,
			Method:   method,
			Password: password,
			PEMFile:  pemPath,
		}

		// Save to vault
		if err := vault.Save(entry); err != nil {
			fmt.Println("❌ Failed to save entry:", err)
			return
		}

		fmt.Println("✅ Saved successfully!")

		// Ask if user wants to test the connection
		fmt.Print("Test connection now? (y/n): ")
		resp, _ := reader.ReadString('\n')
		resp = strings.TrimSpace(strings.ToLower(resp))

		if resp == "y" {
			fmt.Println("🔍 Testing connection...")
			if err := ssh.TestConnection(&entry); err != nil {
				fmt.Println("❌ Connection test failed:", err)
				fmt.Println("💡 You can update or re-add this host anytime using `sshq add`.")
				return
			}
			fmt.Println("✅ Connection successful!")
		} else {
			fmt.Println("ℹ️ Skipped connection test.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

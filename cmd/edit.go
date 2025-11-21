package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"sshq/internal/vault"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit [alias]",
	Short: "Edit a saved SSH connection",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("❌ Provide an alias to edit. Example: sshq edit fyre")
			return
		}

		alias := args[0]
		entries, _ := vault.LoadAll()

		conn, ok := entries[alias]
		if !ok {
			fmt.Printf("❌ Alias '%s' not found\n", alias)
			return
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select field to edit:")
		fmt.Printf("1. Alias: %s\n", conn.Alias)
		fmt.Printf("2. Host: %s\n", conn.Host)
		fmt.Printf("3. User: %s\n", conn.User)
		fmt.Printf("4. Method: %s\n", conn.Method)
		fmt.Printf("5. Password: ****\n")
		fmt.Printf("6. PEM file: %s\n", conn.PEMFile)
		fmt.Print("Enter number: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		fmt.Print("Enter new value: ")
		val, _ := reader.ReadString('\n')
		val = strings.TrimSpace(val)

		oldAlias := conn.Alias

		switch choice {
		case "1":
			conn.Alias = val
		case "2":
			conn.Host = val
		case "3":
			conn.User = val
		case "4":
			conn.Method = val
		case "5":
			conn.Password = val
		case "6":
			conn.PEMFile = val
		default:
			fmt.Println("❌ Invalid choice")
			return
		}

		// handle rename
		if conn.Alias != oldAlias {
			delete(entries, oldAlias)
		}

		entries[conn.Alias] = conn
		vault.SaveAll(entries)
		fmt.Println("✨ Updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
}

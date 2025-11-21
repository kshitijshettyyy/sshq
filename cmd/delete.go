package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sshq/internal/vault"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a saved SSH connection from the vault",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter alias name to delete: ")
		alias, _ := reader.ReadString('\n')
		alias = strings.TrimSpace(alias)

		entries, err := vault.LoadAll()
		if err != nil {
			fmt.Println("❌ Failed to load vault:", err)
			return
		}

		if len(entries) == 0 {
			fmt.Println("ℹ️ Vault is empty — nothing to delete.")
			return
		}

		if _, found := entries[alias]; !found {
			fmt.Println("❌ Alias not found in vault.")
			return
		}

		delete(entries, alias)

		if err := vault.SaveAll(entries); err != nil {
			fmt.Println("❌ Failed to update vault:", err)
			return
		}

		fmt.Printf("🗑️  Found and deleted '%s'\n", alias)
		fmt.Println("✅ Deleted successfully!")
	},
}

func init() { rootCmd.AddCommand(deleteCmd) }

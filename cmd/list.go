package cmd

import (
	"fmt"
	"sshq/internal/vault"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved SSH connections",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := vault.LoadAll()
		if err != nil {
			fmt.Println("❌ Failed to load vault:", err)
			return
		}
		if len(entries) == 0 {
			fmt.Println("No saved connections found.")
			return
		}
		fmt.Println("🗂  Saved Connections:")
		for alias, e := range entries {
			fmt.Printf("• %s (%s@%s, %s)\n", alias, e.User, e.Host, e.Method)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

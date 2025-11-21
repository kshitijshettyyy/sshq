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
		entries, err := vault.List()
		if err != nil {
			fmt.Println("❌", err)
			return
		}
		fmt.Println("🗂  Saved Connections:")
		for _, e := range entries {
			fmt.Printf("• %s (%s@%s, %s)\n", e.Alias, e.User, e.Host, e.Method)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

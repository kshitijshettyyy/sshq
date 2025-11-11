package cmd

import (
	"fmt"
	"sshq/internal/ssh"
	"sshq/internal/vault"

	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect [alias]",
	Short: "Connect to a saved server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		alias := args[0]
		entry, err := vault.Load(alias)
		if err != nil {
			fmt.Println("❌ Failed to load entry:", err)
			return
		}

		fmt.Printf("🔗 Connecting to %s@%s...\n", entry.User, entry.Host)
		if err := ssh.Connect(entry); err != nil {
			fmt.Println("❌ Connection failed:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)
}

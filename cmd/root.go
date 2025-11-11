package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sshq",
	Short: "sshq is a smart SSH connection manager",
	Long:  "sshq helps developers securely store credentials and quickly connect to remote servers.",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// subcommands added in other files
}

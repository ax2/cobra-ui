package cmd

import (
	"github.com/ax2/cobra-ui/examples/devx/cmd/codec"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devx",
	Short: "Developer's utilities, a demo of cobra-ui",
}

func init() {
	rootCmd.AddCommand(codec.GetCommands())
}

func RootCmd() *cobra.Command {
	return rootCmd
}

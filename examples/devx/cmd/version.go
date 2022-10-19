package cmd

import (
	"fmt"

	cobraui "github.com/ax2/cobra-ui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:         "version",
	Short:       "Print the version number of simple test",
	Annotations: cobraui.Options(),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
	},
}

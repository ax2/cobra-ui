package cobraui

import (
	"github.com/ax2/cobra-ui/web"
	"github.com/spf13/cobra"
)

var UICmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the web based UI for cobra based program",
	Run: func(cmd *cobra.Command, args []string) {
		web.StartServer(cmd.Root())
	},
}

func Options(options ...string) map[string]string {
	result := make(map[string]string)
	result["ui"] = "1"
	for _, o := range options {
		result[o] = "1"
	}
	return result
}

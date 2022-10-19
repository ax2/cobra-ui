package cobraui

import (
	"github.com/spf13/cobra"
)

type UiCmdType cobra.Command

var UICmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the web based UI for cobra based program",
	Run: func(cmd *cobra.Command, args []string) {
		startServer(cmd.Root())
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

func (c *UiCmdType) UiEnabled() bool {
	if c.Annotations == nil {
		return false
	}

	if c.Annotations["ui"] == "1" {
		return true
	}

	return false
}

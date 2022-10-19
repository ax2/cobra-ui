package qr

import (
	cobraui "github.com/ax2/cobra-ui"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:         "qr cmd [options]",
	Short:       "qr utilities",
	Annotations: cobraui.Options(),
}

func init() {

}

func GetCommands() *cobra.Command {
	return subCmd
}

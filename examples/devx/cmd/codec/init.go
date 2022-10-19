package codec

import (
	cobraui "github.com/ax2/cobra-ui"
	"github.com/ax2/cobra-ui/examples/devx/cmd/codec/qr"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:         "codec cmd [options]",
	Short:       "codec utilities",
	Annotations: cobraui.Options(),
}

func init() {
	subCmd.AddCommand(qr.GetCommands())
}

func GetCommands() *cobra.Command {
	return subCmd
}

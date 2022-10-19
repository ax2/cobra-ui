package qr

import (
	"fmt"

	cobraui "github.com/ax2/cobra-ui"
	"github.com/spf13/cobra"
)

func init() {
	subCmd.AddCommand(qrEncodeCmd)
	subCmd.AddCommand(qrDecodeCmd)
}

var (
	qrEncodeCmd = &cobra.Command{
		Use:         "encode text",
		Short:       "qr encode",
		Args:        cobra.MinimumNArgs(1),
		Annotations: cobraui.Options(),
		Run:         qrencode,
	}

	qrDecodeCmd = &cobra.Command{
		Use:         "decode text",
		Short:       "qr decode",
		Args:        cobra.MinimumNArgs(1),
		Annotations: cobraui.Options(),
		Run:         qrdecode,
	}
)

func init() {

}

func qrencode(cmd *cobra.Command, args []string) {
	fmt.Println("qrencode...")
}

func qrdecode(cmd *cobra.Command, args []string) {
	fmt.Println("qrdecode...")
}

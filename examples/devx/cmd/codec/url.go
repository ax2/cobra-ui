package codec

import (
	"fmt"
	"net/url"

	cobraui "github.com/ax2/cobra-ui"
	"github.com/spf13/cobra"
)

func init() {
	subCmd.AddCommand(urlEncodeCmd)
	subCmd.AddCommand(urlDecodeCmd)
}

var (
	test         string
	test2        string
	urlEncodeCmd = &cobra.Command{
		Use:         "urlencode text",
		Short:       "url encode",
		Args:        cobra.MinimumNArgs(1),
		Annotations: cobraui.Options(),
		Run:         urlencode,
	}

	urlDecodeCmd = &cobra.Command{
		Use:         "urldecode text",
		Short:       "url decode",
		Args:        cobra.MinimumNArgs(1),
		Annotations: cobraui.Options(),
		Run:         urldecode,
	}
)

func init() {
	urlEncodeCmd.Flags().StringVarP(&test, "test", "", "test", "This is a test flag")
	urlEncodeCmd.Flags().StringVarP(&test2, "test2", "", "test2", "The is another test flag")
}

func urlencode(cmd *cobra.Command, args []string) {
	fmt.Println(url.QueryEscape(args[0]))
	fmt.Printf("test=%s test2=%s\n", test, test2)
}

func urldecode(cmd *cobra.Command, args []string) {
	result, _ := url.QueryUnescape(args[0])
	fmt.Println(result)
}

package main

import (
	"fmt"
	"os"

	cobraui "github.com/ax2/cobra-ui"
	"github.com/ax2/cobra-ui/examples/devx/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()
	rootCmd.AddCommand(cobraui.UICmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

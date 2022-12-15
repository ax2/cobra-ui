package web

import (
	"bytes"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func CmdExecuteAction(c *gin.Context) {
	cmdParam := c.DefaultQuery("cmd", "")
	argsParam := c.DefaultQuery("args", "")
	flagsParam := c.DefaultQuery("flags", "")
	cmds := strings.Split(cmdParam, ".")
	//fmt.Printf("CmdExecuteAction cmd=%s args=%v flags=%v\n", cmdParam, argsParam, flagsParam)
	currentCmd := rootCmd
	var selectedCmd *cobra.Command
	os.Args = []string{"devx"}
	for i := 0; i < len(cmds); i++ {
		for _, cmd := range currentCmd.Commands() {
			if cmd.Name() == cmds[i] {
				os.Args = append(os.Args, cmd.Name())
				if i == len(cmds)-1 {
					selectedCmd = cmd
					break
				} else {
					currentCmd = cmd
				}
			}
		}
	}
	if selectedCmd == nil {
		return
	}
	if argsParam != "" {
		os.Args = append(os.Args, strings.Split(argsParam, " ")...)
		for i := 0; i < len(os.Args); i++ {
			os.Args[i] = strings.Trim(os.Args[i], " ")
		}
	}
	if flagsParam != "" {
		values, err := url.ParseQuery(flagsParam)
		if err == nil {
			for k, v := range values {
				if len(v) > 0 {
					os.Args = append(os.Args, "--"+k) // TODO: short
					os.Args = append(os.Args, v[0])
				}
			}
		}
	}

	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	selectedCmd.SetOut(w)
	os.Stderr = w
	selectedCmd.SetErr(w)

	selectedCmd.Execute()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = oldOut
	//os.Stderr = oldErr
	out := <-outC

	out = strings.Replace(out, "\n", "</br>", -1)
	//fmt.Printf("selectedCmd:%s\nargs:%v\nresult:%s", selectedCmd.Name(), os.Args, out)
	data := map[string]interface{}{
		"result": out,
	}
	c.JSON(0, data)
}

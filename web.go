package cobraui

import (
	"bytes"
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CmdArgInfo struct {
	Type        string
	Name        string
	Description string
	Default     string
}
type CmdInfo struct {
	Name        string
	Description string
	Selected    bool
}

const (
	MAX_CMD_DEEP = 4
)

var (
	rootCmd *cobra.Command

	//go:embed templates
	tmpl embed.FS

	//go:embed static
	static embed.FS
)

func ShowCommand(c *gin.Context) {
	data := map[string]interface{}{}
	c.HTML(http.StatusOK, "index.tmpl", data)
}

func CmdExecuteAction(c *gin.Context) {
	cmdParam := c.DefaultQuery("cmd", "")
	argsParam := c.DefaultQuery("args", "")
	flagsParam := c.DefaultQuery("flags", "")
	cmds := strings.Split(cmdParam, ".")
	fmt.Printf("CmdExecuteAction cmd=%s args=%v flags=%v\n", cmdParam, argsParam, flagsParam)
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

	//bufOut := new(bytes.Buffer)
	//selectedCmd.SetOut(bufOut)

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	selectedCmd.SetOut(w)

	selectedCmd.Execute()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	w.Close()
	os.Stdout = old
	out := <-outC

	fmt.Printf("selectedCmd:%s\nargs:%v\nresult:%s", selectedCmd.Name(), os.Args, out)
	data := map[string]interface{}{
		"result": out,
	}
	c.JSON(0, data)
}

func CmdInfoAction(c *gin.Context) {
	cmdParam := c.DefaultQuery("cmd", "")
	cmdInfos := make([][]*CmdInfo, MAX_CMD_DEEP)
	if rootCmd == nil {
		return
	}

	cmds := strings.Split(cmdParam, ".")
	if cmdParam == "" {
		cmds = []string{}
	}
	fmt.Printf("cmd=%+v\n", cmds)

	currentCmd := rootCmd
	for i := 0; i < MAX_CMD_DEEP; i++ {
		var selectedCmd *cobra.Command
		selectedCmd = nil
		for j, cmd := range currentCmd.Commands() {
			if !cmd.IsAvailableCommand() || cmd.IsAdditionalHelpTopicCommand() {
				continue
			}
			cc := (*UiCmdType)(cmd)
			if !cc.UiEnabled() {
				continue
			}
			selected := false
			if i < len(cmds) {
				if cmd.Name() == cmds[i] {
					selectedCmd = cmd
					selected = true
				}
			} else {
				if i >= len(cmds) && j == 0 {
					selectedCmd = cmd
					selected = true
				}
			}
			cmdInfos[i] = append(cmdInfos[i], &CmdInfo{
				Name:     cmd.Name(),
				Selected: selected,
			})
		}
		if selectedCmd != nil {
			currentCmd = selectedCmd
		} else {
			break
		}
	}

	data := map[string]interface{}{
		"cmds":        cmdInfos,
		"use":         currentCmd.Use,
		"description": currentCmd.Short,
		"flags":       parseFlags(currentCmd.Flags()),
	}
	c.JSON(0, data)
}

func parseFlags(f *pflag.FlagSet) (result []*CmdArgInfo) {
	f.VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("- %+v\n", flag)
		if flag.Hidden || len(flag.Deprecated) != 0 {
			return
		}

		_, usage := pflag.UnquoteUsage(flag)
		result = append(result, &CmdArgInfo{
			Name:        flag.Name,
			Description: usage,
			Type:        flag.Value.Type(),
		})
	})
	fmt.Printf("result=%v\n", result)

	return
}

func startServer(cmd *cobra.Command) {
	rootCmd = cmd
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	t, _ := template.ParseFS(tmpl, "templates/*.tmpl")
	r.SetHTMLTemplate(t)
	r.Any("/static/*filepath", func(c *gin.Context) {
		staticServer := http.FileServer(http.FS(static))
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

	r.Any("/", ShowCommand)
	r.GET("/cmdinfo", CmdInfoAction)
	r.GET("/execute", CmdExecuteAction)

	_ = r.Run(":8189")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
}

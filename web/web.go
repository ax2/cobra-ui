package web

import (
	"embed"
	_ "embed"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type UiCmdType cobra.Command

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

func (c *UiCmdType) UiEnabled() bool {
	if c.Annotations == nil {
		return false
	}

	if c.Annotations["ui"] == "1" {
		return true
	}

	return false
}

func ShowCommand(c *gin.Context) {
	data := map[string]interface{}{}
	c.HTML(http.StatusOK, "index.tmpl", data)
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

func StartServer(cmd *cobra.Command) {
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

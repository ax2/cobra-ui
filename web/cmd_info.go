package web

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

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

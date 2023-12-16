package controller

import (
	"fmt"
	"glugin/utils"
	"time"

	"github.com/martinlindhe/notify"
	"github.com/neovim/go-client/nvim/plugin"
	"github.com/xanzy/go-gitlab"
)

const NeovimStr = "neovim"
const PluginName = "mypluin"

type Controller struct {
	*plugin.Plugin
	gitlabCli *gitlab.Client
}

func NewController(p *plugin.Plugin) *Controller {
	ctrl := &Controller{
		Plugin: p,
	}
	return ctrl
}

func (ctrl *Controller) Serve() error {
	ctrl.HandleAutocmd(&plugin.AutocmdOptions{
		Event:   "VimEnter",
		Group:   "myplugin",
		Pattern: "*",
	}, func() {
		ctrl.startBackGroupJobs()
	})

	ctrl.HandleFunction(&plugin.FunctionOptions{Name: "RpcPing"}, ctrl.rpcPing)
	ctrl.HandleFunction(&plugin.FunctionOptions{Name: "ScheduleTask"}, ctrl.scheduleTask)
	return nil
}

func (ctrl *Controller) rpcPing(args []string) ([]string, error) {
	defer func() {
		if e := recover(); e != nil {
			notify.Alert(NeovimStr, PluginName, fmt.Sprintf("panic: %v", utils.Marshal(e)), "")
		}
	}()
	/*
		content, _ := ctrl.getCursorFunction()
		ctrl.nvimNotify(PluginName, fmt.Sprintf("```go\n%v\n```", content))
	*/
	return []string{"2", "1"}, nil
}

func (ctrl *Controller) scheduleTask(args []string) (string, error) {
	ctrl.Nvim.ExecLua(`print(tostring(vim.oxi.schedule_task()))`, nil)
	return "", nil
}

func (ctrl *Controller) startBackGroupJobs() {
	go func() {
		for range time.Tick(time.Second * 10) {
			// notify.Notify("neovim", "", "该休息了哦", "")
			/*
				pos, err := ctrl.getPosition()
				if err != nil {
					ctrl.nvimNotify(PluginName, "getPosition err: %v", err)
					continue
				}
				ctrl.nvimNotify(PluginName, ctrl.Marshal(pos))
			*/
		}
	}()
}

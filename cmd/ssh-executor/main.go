package main

import (
	"os"
	"strings"
	"fmt"

	"git.eju-inc.com/ops/go-common/log"
	"git.eju-inc.com/ops/go-common/log/ctx"
	"git.eju-inc.com/ops/go-common/version"

	"github.com/bypdhu/ssh-executor/cmd/ssh-executor/flags"
	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/module"
	"github.com/bypdhu/ssh-executor/common"
	"github.com/bypdhu/ssh-executor/task"
)

var (
	c *conf.Config
	cs map[string]*conf.Config
	//rs map[string]*result.DirectResult
)

func main() {
	fs := flags.ParseFlags(os.Args)
	log.SetContextExtractor(ctx.ZipkinTraceExtractor{})

	c = conf.Load(fs.ConfigFilePath)
	//log.Infof("config from file and default: %s", c)

	flags.OverrideConfWithFlags(c, fs)
	//log.Infof("final config, override by command-args: %s", c)
	//log.Infof("==========launch type: %s\n", c.LaunchType)

	switch c.LaunchType {
	case common.LAUNCH_DIRECT.String(), strings.ToLower(common.LAUNCH_DIRECT.String()):
		doDirect(c)
	case common.LAUNCH_SERVER.String(), strings.ToLower(common.LAUNCH_SERVER.String()):
		log.Infof("Starting up %s ...\n", version.Print())
		log.Infoln(fs.PrintAll())
		log.Infof("Final config, override by command-args: %s", c)

		doServer(c)
	}
}

func doServer(c *conf.Config) {

}

func doDirect(c *conf.Config) {

	addTaskFromCmd(c)

	hs := conf.GetHosts(c)

	cs = conf.GetCopiedConfigMap(c, hs)

	module.RunAll(cs)

	for h, c := range cs {
		fmt.Printf("+++++++host:%s\n", h)
		for _, t := range c.Tasks {
			fmt.Printf("task:%+v\n", t)
		}
	}
}

func addTaskFromCmd(c *conf.Config) {
	_t := task.DefaultTask(c.Direct.Module)

	switch c.Direct.Module {
	case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
		if c.Direct.Command != "" {
			_t.Command = c.Direct.Command
			c.Tasks = append(c.Tasks, _t)
			copy(c.Tasks[1:], c.Tasks[0:len(c.Tasks) - 1])
			c.Tasks[0] = _t
		}
	case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
	//	TODO
	default:
		if c.Direct.Command != "" {
			_t.Command = c.Direct.Command
			c.Tasks = append(c.Tasks, _t)
			copy(c.Tasks[1:], c.Tasks[0:len(c.Tasks) - 1])
			c.Tasks[0] = _t		}
	}
}


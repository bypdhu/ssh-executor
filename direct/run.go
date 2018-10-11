package direct

import (
	"fmt"
	"strings"

	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/module"
	"github.com/bypdhu/ssh-executor/task"
	"github.com/bypdhu/ssh-executor/common"
)

var (
	cs map[string]*conf.Config
)

func Run(c *conf.Config) {

	addTaskFromCmd(c)

	hs := conf.GetHostsFromConfig(c)

	cs = conf.GetCopiedConfigMap(c, hs)

	module.RunAll(cs)

	//for h, c := range cs {
	//	fmt.Printf("+++++++host:%s\n", h)
	//	for _, t := range c.Tasks {
	//		fmt.Printf("task:%+v\n", t)
	//	}
	//}

	rs := conf.GenerateResultToString(cs, "json")
	fmt.Printf("%s\n", rs)
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
			c.Tasks[0] = _t
		}
	}
}


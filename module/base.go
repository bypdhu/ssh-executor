package module

import (
	"sync"
	"strings"

	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/module/shell"
	"github.com/bypdhu/ssh-executor/common"
	"github.com/bypdhu/ssh-executor/task"
)

var (
	wg sync.WaitGroup
)

func RunAll(cs map[string]*conf.Config) {

	for _, c := range cs {
		wg.Add(1)
		go runTasks(c, &wg)
	}
	wg.Wait()

}

func runTasks(c *conf.Config, w *sync.WaitGroup) {
	defer w.Done()
	for _, t := range c.Tasks {
		t.HostDup = c.HostDup
		switch t.Module  {
		case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
			runShell(c, t)
		case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
			runCopy(c, t)
		default:
			runShell(c, t)
		}
	}
}

func runShell(c *conf.Config, t *task.Task) {
	shell.Run(c, t)
}

func runCopy(c *conf.Config, t *task.Task) {
}


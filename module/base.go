package module

import (
	"sync"

	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/module/shell"
)

const (
	MODULE_SHELL = "shell"
	MODULE_COPY = "copy"
)

var (
	wg sync.WaitGroup
)

func RunAll(c *conf.Config, hs []string, r map[string]*result.SSHResult) {

	for _, h := range hs {
		r[h] = &result.SSHResult{}

		wg.Add(1)
		switch c.Direct.Module {
		case MODULE_SHELL:
			go runShell(c, h, r[h], &wg)
		case MODULE_COPY:
			go runCopy(c, h, r[h], &wg)
		default:
			go runShell(c, h, r[h], &wg)
		}
	}
	wg.Wait()

}

func runShell(c *conf.Config, h string, r *result.SSHResult, w *sync.WaitGroup) {
	defer w.Done()
	shell.Run(c, h, r)
}

func runCopy(c *conf.Config, h string, r *result.SSHResult, w *sync.WaitGroup) {
	defer w.Done()
}
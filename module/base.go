package module

import (
	"sync"

	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/module/shell"
	"github.com/bypdhu/ssh-executor/common"
)

var (
	wg sync.WaitGroup
)

func RunAll(c *conf.Config, hs []string, r map[string]*result.BaseResult) {

	for _, h := range hs {
		r[h] = &result.BaseResult{}

		wg.Add(1)
		switch c.Direct.Module {
		case common.MODULE_SHELL.String():
			go runShell(c, h, r[h], &wg)
		case common.MODULE_COPY.String():
			go runCopy(c, h, r[h], &wg)
		default:
			go runShell(c, h, r[h], &wg)
		}
	}
	wg.Wait()

}

func runShell(c *conf.Config, h string, r *result.BaseResult, w *sync.WaitGroup) {
	defer w.Done()
	shell.Run(c, h, r)
}

func runCopy(c *conf.Config, h string, r *result.BaseResult, w *sync.WaitGroup) {
	defer w.Done()
}
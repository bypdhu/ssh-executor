package result

import (
	"github.com/bypdhu/ssh-executor/bssh"
)

type Result struct {
	Result  string
	Err     error
	Success bool
}

type SSHResult struct {
	bssh.PerRun
	Err error
}


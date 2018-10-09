package result

import "github.com/bypdhu/ssh-executor/common"

type BaseResult struct {
	SSHResult
	SFTPResult
	Err error
}

type SSHResult struct {
	Result   string
	ExitCode int
}

type SFTPResult struct {
	Changed bool
}

type SSHOneResult struct {
	SSHResult
	Id int
}

type SFTPOneResult struct {
	SFTPResult
	Id int
}

type OneResult struct {
	BaseResult
	Id int
}

type DirectResult struct {
	Module  common.ModuleType

	SSHRes  SSHResult
	SFTPRes []*SFTPOneResult
}


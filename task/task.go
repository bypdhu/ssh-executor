package task

import (
	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/common"
)

type Task struct {
	Args
	Module common.ModuleType
	Name   string
}

type Args struct {
	ShellArgs
	CopyArgs
}

type ShellArgs struct {
	Sudo
	result.BaseResult

	Command         string
	OriginalCommand string
	Chdir           string
	Login           bool
}

type Sudo struct {
	Become       bool
	BecomeMethod string
	BecomeUser   string
}

type CopyArgs struct {
	SftpMode  common.SftpMode
	CopyFiles []*CopyOneFile
	IgnoreErr bool // if true, will continue run when copy many files.
}

type CopyOneFile struct {
	Sudo
	result.BaseResult

	Src             string
	Dest            string
	Owner           string
	Group           string
	//Mode            os.FileMode
	Mode            string
	Md5             string
	ForceCopy       bool

	CreateDirectory bool
	Recursive       bool
	DirectoryMode   string
}

func DefaultTask(m common.ModuleType) *Task {
	switch m {
	case common.MODULE_SHELL:
		return &Task{
			Module:common.MODULE_SHELL,
			Name:"shell task",
			Args:Args{
				ShellArgs:ShellArgs{
					Chdir:"",
					Login:true,
					Sudo:Sudo{Become:false, BecomeMethod:"sudo", BecomeUser:"root"},
					Command:"",
					OriginalCommand:"",
				}},
		}
	case common.MODULE_COPY:
		return &Task{
			Module:common.MODULE_COPY,
			Name:"copy task",
			Args:Args{
				CopyArgs:CopyArgs{
					CopyFiles:[]*CopyOneFile{{Src:"", Dest:""}}},
			},
		}
	default:
		return &Task{
			Module:common.MODULE_SHELL,
			Name:"shell task",
			Args:Args{
				ShellArgs:ShellArgs{
					Login:true,
					Sudo:Sudo{Become:false, BecomeMethod:"sudo", BecomeUser:"root"},
				}},
		}
	}
}
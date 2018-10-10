package task

import (
	"strings"

	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/common"
)

type Task struct {
	Args            `yaml:"args"`
	Module  string  `yaml:"module"`
	Name    string  `yaml:"name"`

	HostDup string
}

type Args struct {
	ShellArgs  `yaml:"shell"`
	CopyArgs   `yaml:"copy"`
}

type ShellArgs struct {
	Sudo
	result.BaseResult

	OriginalCommand string

	Command         string  `yaml:"command"`
	Chdir           string  `yaml:"chdir"`
	Login           bool    `yaml:"login"`
}

type Sudo struct {
	Become       bool    `yaml:"become"`
	BecomeMethod string  `yaml:"become_method"`
	BecomeUser   string  `yaml:"become_user"`
}

type CopyArgs struct {
	Sudo
	SftpMode  string          `yaml:"sftp_mode"`
	CopyFiles []*CopyOneFile  `yaml:"copy_files"`
	IgnoreErr bool            `yaml:"ignore_err"` // if true, will continue run when copy many files.
}

type CopyOneFile struct {
	result.BaseResult

	Src             string  `yaml:"src"`
	Dest            string  `yaml:"dest"`
	Owner           string  `yaml:"owner"`
	Group           string  `yaml:"group"`
	//Mode            os.FileMode
	Mode            string  `yaml:"mode"`
	Md5             string  `yaml:"md5"`
	ForceCopy       bool    `yaml:"force"`

	CreateDirectory bool    `yaml:"create_directory"`
	Recursive       bool    `yaml:"recursive"`
	DirectoryMode   string  `yaml:"directory_mode"`
}

func DefaultTask(m string) *Task {
	switch m {
	case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
		return &Task{
			Module:common.MODULE_SHELL.String(),
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
	case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
		return &Task{
			Module:common.MODULE_COPY.String(),
			Name:"copy task",
			Args:Args{
				CopyArgs:CopyArgs{
					CopyFiles:[]*CopyOneFile{{Src:"", Dest:""}}},
			},
		}
	default:
		return &Task{
			Module:common.MODULE_SHELL.String(),
			Name:"shell task",
			Args:Args{
				ShellArgs:ShellArgs{
					Login:true,
					Sudo:Sudo{Become:false, BecomeMethod:"sudo", BecomeUser:"root"},
				}},
		}
	}
}
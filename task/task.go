package task

import (
    "strings"

    "github.com/bypdhu/ssh-executor/result"
    "github.com/bypdhu/ssh-executor/common"
)

type Task struct {
    Args            `yaml:"args" json:"args"`
    Module  string  `yaml:"module" json:"module"`
    Name    string  `yaml:"name" json:"name"`

    HostDup string
}

type Args struct {
    ShellArgs  `yaml:"shell" json:"shell"`
    CopyArgs   `yaml:"copy" json:"copy"`
}

type ShellArgs struct {
    Sudo
    result.BaseResult

    OriginalCommand string

    Command         string  `yaml:"command" json:"command"`
    Chdir           string  `yaml:"chdir" json:"chdir"`
    Login           bool    `yaml:"login" json:"login"`
}

type Sudo struct {
    Become       bool    `yaml:"become" json:"become"`
    BecomeMethod string  `yaml:"become_method" json:"become_method"`
    BecomeUser   string  `yaml:"become_user" json:"become_user"`
}

type CopyArgs struct {
    Sudo
    SftpMode  string          `yaml:"sftp_mode" json:"sftp_mode"`
    CopyFiles []*CopyOneFile  `yaml:"copy_files" json:"copy_files"`
    IgnoreErr bool            `yaml:"ignore_err" json:"ignore_err"` // if true, will continue run when copy many files.
}

type CopyOneFile struct {
    result.BaseResult

    Src             string  `yaml:"src" json:"src"`
    Dest            string  `yaml:"dest" json:"dest"`
    Owner           string  `yaml:"owner" json:"owner"`
    Group           string  `yaml:"group" json:"group"`
    //Mode            os.FileMode
    Mode            string  `yaml:"mode" json:"mode"`
    Md5             string  `yaml:"md5" json:"md5"`
    ForceCopy       bool    `yaml:"force" json:"force"`

    CreateDirectory bool    `yaml:"create_directory" json:"create_directory"`
    Recursive       bool    `yaml:"recursive" json:"recursive"`
    DirectoryMode   string  `yaml:"directory_mode" json:"directory_mode"`
}

func DefaultTask(m string) *Task {
    switch m {
    case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
        return &Task{
            Module:common.MODULE_SHELL.String(),
            Name:"shell task",
            Args:Args{
                ShellArgs:ShellArgs{
                    Login:true,
                    Sudo:Sudo{Become:false, BecomeMethod:"sudo", BecomeUser:"root"},
                }},
        }
    case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
        return &Task{
            Module:common.MODULE_COPY.String(),
            Name:"copy task",
            //Args:Args{
            //	CopyArgs:CopyArgs{
            //		CopyFiles:[]*CopyOneFile{}},
            //},
        }
    default:
        return &Task{
            Module:common.MODULE_SHELL.String(),
            Name:"shell default task",
            Args:Args{
                ShellArgs:ShellArgs{
                    Login:true,
                    Sudo:Sudo{Become:false, BecomeMethod:"sudo", BecomeUser:"root"},
                }},
        }
    }
}
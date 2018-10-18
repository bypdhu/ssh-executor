package conf

import (
    "testing"
    "fmt"

    "github.com/bypdhu/ssh-executor/utils"
    "github.com/bypdhu/ssh-executor/task"
    "github.com/bypdhu/ssh-executor/result"
)

func TestLoad(t *testing.T) {
    cp, _ := utils.GetFullPath(".")
    c := Load(cp + "/../testdata/config.yml")
    fmt.Printf("%s", c)
    fmt.Printf("%s", c.original)
}

func TestGetCopiedConfigMap(t *testing.T) {
    c1 := &Config{
        Tasks:[]*task.Task{
            {Module:"copy", Name:"name1", HostDup:"1.1.1.1",
                Args:task.Args{CopyArgs:task.CopyArgs{SftpMode:"push", CopyFiles:[]*task.CopyOneFile{
                    {Src:"src1", Dest:"dest1", BaseResult:result.BaseResult{SFTPResult:result.SFTPResult{Changed:true}}},
                    {Src:"src2", Dest:"dest2", BaseResult:result.BaseResult{SFTPResult:result.SFTPResult{Changed:true}}},
                }}}},
            {Module:"shell", Name:"name2", HostDup:"2.2.2.2",
                Args:task.Args{ShellArgs:task.ShellArgs{Command:"cmd1", BaseResult:result.BaseResult{SSHResult:result.SSHResult{Result:"result1", ExitCode:0}}}},
            }},
        SSHConfig:SSHConfig{Timeout:12},
    }
    hs := []string{"1.1", "1.2", "1.3"}

    cs := GetCopiedConfigMap(c1, hs)

    for h, c := range cs {
        fmt.Printf("%s,%+v\n", h, c)
        for _, t := range c.Tasks {
            fmt.Printf("%+v\n", t)
        }
    }

    fmt.Println()

    cs["1.1"].LaunchType = "launch1"
    cs["1.2"].LaunchType = "launch2"
    cs["1.3"].Tasks[0].Args.CopyArgs.Sudo.BecomeUser = "user.."
    cs["1.3"].Tasks[0].Args.CopyArgs.CopyFiles[0].Src = "src___changed"
    for h, c := range cs {
        fmt.Printf("%s,%+v\n", h, c)
        for _, t := range c.Tasks {
            fmt.Printf("%+v\n", t)
            if len(t.CopyFiles) > 0 {
                fmt.Printf("%+v\n", t.CopyFiles[0])
            }
        }
    }
}
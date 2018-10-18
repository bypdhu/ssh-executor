package task

import (
    "testing"
    "fmt"
    "github.com/pkg/errors"

    "github.com/bypdhu/ssh-executor/common"
    "github.com/bypdhu/ssh-executor/result"
)

func TestTask_TaskResultToJsonString(t *testing.T) {
    a := DefaultTask(common.MODULE_SHELL.String())
    fmt.Printf("%+v\n", a)
    a.ShellArgs.Err = errors.New("this is err")
    fmt.Printf("%s\n", TaskResultToJsonString(a))
    a.ShellArgs.ExitCode = 1
    fmt.Printf("%s\n", TaskResultToJsonString(a))

    fmt.Println()

    b := DefaultTask(common.MODULE_COPY.String())
    fmt.Printf("%s\n", TaskResultToJsonString(b))

    b.SftpMode = "pull"
    b.CopyFiles = []*CopyOneFile{
        {Src:"src1", Dest:"dest1", BaseResult:result.BaseResult{SFTPResult:result.SFTPResult{Changed:false}}},
        {Src:"src2", Dest:"dest2", BaseResult:result.BaseResult{SFTPResult:result.SFTPResult{Changed:false}}},
    }
    fmt.Printf("%+v\n", b)
    fmt.Printf("%s\n", TaskResultToJsonString(b))

    b.CopyFiles[0].Err = errors.New("this is err")
    fmt.Printf("%s\n", TaskResultToJsonString(b))

}
package task

import (
    "strings"
    "encoding/json"

    "github.com/bypdhu/ssh-executor/common"
    "github.com/bypdhu/ssh-executor/result"
)

func TaskResultToJson(t *Task) ([]byte, error) {
    switch t.Module {
    case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
        success := t.ShellArgs.SSHResult.ExitCode == 0 && t.BaseResult.Err == nil
        e := ""
        if t.Err != nil {
            e = t.Err.Error()
        }
        return json.Marshal(result.ShellTaskResult{
            Name:t.Name, Module:t.Module, Result:t.SSHResult, Success:success, Err:e})

    case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
        copyFiles := []result.CopyOneFileResult{}
        success := true
        for _, i := range t.CopyFiles {
            e := ""
            if i.BaseResult.Err != nil {
                e = i.Err.Error()
                success = false
            }

            copyFile := result.CopyOneFileResult{Src:i.Src, Dest:i.Dest, Result:i.SFTPResult, Err:e}
            copyFiles = append(copyFiles, copyFile)
        }
        return json.Marshal(result.CopyTaskResult{
            Name:t.Name, Module:t.Module, SftpMode:t.SftpMode, Success:success, CopyFiles:copyFiles})

    default:
        success := t.ShellArgs.SSHResult.ExitCode == 0 && t.BaseResult.Err == nil
        e := ""
        if t.Err != nil {
            e = t.Err.Error()
        }
        return json.Marshal(result.ShellTaskResult{
            Name:t.Name, Module:t.Module, Result:t.SSHResult, Success:success, Err:e})

    }

}

// now no use.
func TaskResultToJsonString(t *Task) string {
    s, err := TaskResultToJson(t)
    if err != nil {
        return ""
    }
    return string(s)
}
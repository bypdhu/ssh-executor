package conf

import (
	"encoding/json"
	"strings"

	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/common"
)

func GenerateResultToString(cs map[string]*Config, t string) string {
	switch t {
	case "json":
		return GenerateResultToJsonString(cs)
	case "raw":
		return ""
	default:
		return GenerateResultToJsonString(cs)
	}

}

func GenerateResultToJsonString(cs map[string]*Config) string {
	s, err := GenerateResultToJson(cs)
	if err != nil {
		return ""
	}
	return string(s)
}

func GenerateResultToJson(cs map[string]*Config) ([]byte, error) {
	return json.Marshal(GenerateResult(cs))
}

func GenerateResult(cs map[string]*Config) result.Result {
	results := result.Result{Success:true}
	hostResults := []*result.HostResult{}

	for h, c := range cs {
		hr := &result.HostResult{Success:true}
		hr.Host = h
		for _, t := range c.Tasks {
			switch t.Module {
			case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
				s := true
				e := ""
				if t.ShellArgs.SSHResult.ExitCode != 0 || t.BaseResult.Err != nil {
					s = false
					hr.Success = false
					results.Success = false
					if t.BaseResult.Err != nil {
						e = t.BaseResult.Err.Error()
					}
				}

				tr := result.ShellTaskResult{
					Name:t.Name, Module:t.Module, Result:result.SSHResult{
						Stdout:t.SSHResult.Stdout, ExitCode:t.SSHResult.ExitCode,
					}, Err:e, Success:s}
				hr.Tasks = append(hr.Tasks, tr)
			case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
				cr := []result.CopyOneFileResult{}
				s := true
				for _, i := range t.CopyFiles {
					e := ""
					if i.BaseResult.Err != nil {
						e = i.Err.Error()
						s = false
						hr.Success = false
						results.Success = false
					}
					copyFile := result.CopyOneFileResult{Src:i.Src, Dest:i.Dest, Result:i.SFTPResult, Err:e}
					cr = append(cr, copyFile)
				}
				tr := result.CopyTaskResult{
					Name:t.Name, Module:t.Module, Success:s, SftpMode:t.SftpMode, CopyFiles:cr,
				}
				hr.Tasks = append(hr.Tasks, tr)
			default:
				s := true
				e := ""
				if t.ShellArgs.SSHResult.ExitCode != 0 || t.BaseResult.Err != nil {
					s = false
					hr.Success = false
					results.Success = false
					if t.BaseResult.Err != nil {
						e = t.BaseResult.Err.Error()
					}
				}

				tr := result.ShellTaskResult{
					Name:t.Name, Module:t.Module, Result:result.SSHResult{
						Stdout:t.SSHResult.Stdout, ExitCode:t.SSHResult.ExitCode,
					}, Err:e, Success:s}
				hr.Tasks = append(hr.Tasks, tr)
			}
		}

		hostResults = append(hostResults, hr)

	}

	results.Detail = hostResults
	return results
}
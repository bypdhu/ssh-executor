package result

import (
	"encoding/json"
	"github.com/bypdhu/ssh-executor/common"
	"strings"
)

func (r *SSHResult) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *SSHResult) ToJsonString() string {
	s, err := r.ToJson()
	if err != nil {
		return ""
	}
	return string(s)
}

func (r *SFTPResult) ToJson() ([]byte, error) {
	return json.Marshal(r)
}

func (r *SFTPResult) ToJsonString() string {
	s, err := r.ToJson()
	if err != nil {
		return ""
	}
	return string(s)
}

func (r *BaseResult)ToJson(m string) ([]byte, error) {
	e := ""
	if r.Err != nil {
		e = r.Err.Error()
	}

	switch m {
	case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
		return json.Marshal(struct {
			Result   string
			ExitCode int
			Err      string
		}{Result:r.Result, ExitCode:r.ExitCode, Err:e})
	case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
		return json.Marshal(struct {
			Changed bool
			Err     string
		}{Changed:r.Changed, Err:e})
	default:
		return json.Marshal(struct {
			Result   string
			ExitCode int
			Err      string
		}{Result:r.Result, ExitCode:r.ExitCode, Err:e})
	}
}

func (r *BaseResult) ToJsonString(m string) string {
	s, err := r.ToJson(m)
	if err != nil {
		return ""
	}
	return string(s)
}
package result

import "encoding/json"

func (r *SSHResult) ToJson() ([]byte, error) {
	e := ""
	if r.Err != nil {
		e = r.Err.Error()
	}

	return json.Marshal(struct {
		Result   string
		ExitCode int
		Err      string
	}{Result:r.Result, ExitCode:r.ExitCode, Err:e})
}

func (r *SSHResult) ToJsonString() string {
	s, err := r.ToJson()
	if err != nil {
		return ""
	}
	return string(s)
}

func (r *SFTPResult) ToJson() ([]byte, error) {
	e := ""
	if r.Err != nil {
		e = r.Err.Error()
	}
	return json.Marshal(struct {
		Changed bool
		Err     string
	}{Changed:r.Changed, Err:e})
}

func (r *SFTPResult)ToJsonString() string {
	s, err := r.ToJson()
	if err != nil {
		return ""
	}
	return string(s)
}
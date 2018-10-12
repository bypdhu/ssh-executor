package result

type Result struct {
	Success bool           `json:"success"`
	Msg     string         `json:"msg"`
	Detail  []*HostResult  `json:"detail"`
}

type HostResult struct {
	Host    string        `json:"host"`
	Success bool          `json:"success"`
	Tasks   []TaskResult  `json:"tasks"`
}

type TaskResult interface {
	taskResult()
}

func (r ShellTaskResult) taskResult() {}
func (r CopyTaskResult) taskResult() {}

type ShellTaskResult struct {
	Name    string     `json:"name"`
	Module  string     `json:"module"`
	Success bool       `json:"success"`
	Err     string     `json:"err"`
	Result  SSHResult  `json:"result"`
}

type CopyTaskResult struct {
	Name      string               `json:"name"`
	Module    string               `json:"module"`
	SftpMode  string               `json:"sftpmode"`
	Success   bool                 `json:"success"`
	Err       string               `json:"err"`
	CopyFiles []CopyOneFileResult  `json:"copyfiles"`
}

type CopyOneFileResult struct {
	Src    string      `json:"src"`
	Dest   string      `json:"dest"`
	Err    string      `json:"err"`
	Result SFTPResult  `json:"result"`
}

type BaseResult struct {
	SSHResult
	SFTPResult
	Err error
}

type SSHResult struct {
	Stdout   string  `json:"stdout"`
	ExitCode int     `json:"exitcode"`
}

type SFTPResult struct {
	Changed bool  `json:"changed"`
}

//type SSHOneResult struct {
//	SSHResult
//	Id int
//}
//
//type SFTPOneResult struct {
//	SFTPResult
//	Id int
//}

//type OneResult struct {
//	BaseResult
//	Id int
//}
//
//type DirectResult struct {
//	Module  string
//
//	SSHRes  SSHResult
//	SFTPRes []*SFTPOneResult
//}


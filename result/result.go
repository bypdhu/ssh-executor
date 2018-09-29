package result

type BaseResult struct {
	SSHResult
	SFTPResult
	Err error
}

type SSHResult struct {
	Result   string
	ExitCode int
}

type SFTPResult struct {
	Changed bool
}
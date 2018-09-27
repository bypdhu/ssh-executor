package result

type Result struct {
	Result  string
	Err     error
	Success bool
}

type SSHResult struct {
	Result   string
	ExitCode int
	Err      error
}

type SFTPResult struct {
	Changed bool
	Err     error
}
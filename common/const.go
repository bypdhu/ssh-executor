package common

type LaunchType string
type ModuleType string
type SftpMode string

const (
	LAUNCH_SERVER LaunchType = "server"
	LAUNCH_DIRECT LaunchType = "direct"

	MODULE_SHELL ModuleType = "SHELL"
	MODULE_COPY ModuleType = "COPY"

	SFTP_PULL SftpMode = "PULL"
	SFTP_PUSH SftpMode = "PUSH"
)

func (l LaunchType) String() string {
	return string(l)
}

func (m ModuleType) String() string {
	return string(m)
}

func (s SftpMode) String() string {
	return string(s)
}
package sftp

import (
	"github.com/bypdhu/ssh-executor/task"
	"github.com/bypdhu/ssh-executor/bssh"
	"github.com/bypdhu/ssh-executor/conf"
)

func Run(c *conf.Config, t *task.Task) {
	client := bssh.NewSftp(t.HostDup, 22, c.SSHConfig.UserName, c.SSHConfig.Password)

	client.Task = t

	client.SftpStart()
}
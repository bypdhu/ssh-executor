package sftp

import (
    "github.com/bypdhu/ssh-executor/task"
    "github.com/bypdhu/ssh-executor/bssh"
    "github.com/bypdhu/ssh-executor/conf"
)

func Run(c *conf.Config, t *task.Task) {
    client, err := bssh.NewSftp(t.HostDup, 22, c.SSHConfig.UserName, c.SSHConfig.Password)

    if err != nil {
        t.Err = err
        return
    }

    client.Task = t

    client.SftpStart()
}
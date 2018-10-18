package shell

import (
    "github.com/bypdhu/ssh-executor/bssh"
    "github.com/bypdhu/ssh-executor/conf"
    "github.com/bypdhu/ssh-executor/task"
)

func Run(c *conf.Config, t *task.Task) {

    client := bssh.New(t.HostDup, 22, c.SSHConfig.UserName, c.SSHConfig.Password)
    //log.Infof("+++++++++now run %s on host %s\n", c.Direct.Command, h)
    //log.Infof("client is %s", client)
    client.Task = t
    client.RunCommandTask()

    //log.Infof("++++++result is %s on host %s\n", result.result, h)
}

package shell

import (
	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/bssh"
	"github.com/bypdhu/ssh-executor/conf"
)

func Run(c *conf.Config, h string, result *result.SSHResult) {

	client := bssh.New(h, 22, c.Direct.UserName, c.Direct.Password)
	//log.Infof("+++++++++now run %s on host %s\n", c.Direct.Command, h)
	//log.Infof("client is %s", client)
	result.Err = client.Run(c.Direct.Command)
	result.Result = client.Session.Result
	result.ExitCode = client.Session.ExitCode

	//log.Infof("++++++result is %s on host %s\n", result.result, h)
}

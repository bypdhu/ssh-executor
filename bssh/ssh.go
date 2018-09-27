package bssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"

	"github.com/bypdhu/ssh-executor/result"
)

type SSHCli struct {
	Cli
	Session Session
}

type Session struct {
	SSHCommand
	session *ssh.Session
}

type SSHCommand struct {
	result.SSHResult
	Cmd     string
	Comment string
	Args    SSHArgs
}

type SSHArgs struct {

}

func New(ip string, port int, username string, password string) *SSHCli {
	cli := new(SSHCli)
	cli.Host = ip
	cli.Port = port
	cli.Username = username
	cli.Password = password
	return cli
}

func (c *SSHCli) Run(cmd string) (err error) {
	//fmt.Printf("client:%s\n", c.client)
	if c.client == nil {
		if err = c.newClient(); err != nil {
			return
		}
	}
	//fmt.Printf("session:%s\n", c.session)
	if c.Session.session == nil {
		if err = c.newSession(); err != nil {
			return
		}
	}

	defer c.closeSession()

	c.Session.ExitCode = 0
	c.Session.Cmd = cmd
	buf, err := c.Session.session.CombinedOutput(cmd)
	if err != nil {
		if v, ok := err.(*ssh.ExitError); ok {
			c.Session.ExitCode = v.Waitmsg.ExitStatus()
		}
		c.Session.ExitCode = -1
	}
	c.Session.Result = fmt.Sprintf("%s", buf)
	//c.Session.LastResult = string(buf)

	return
}

func (c *SSHCli) newSession() (err error) {
	c.Session.session, err = c.client.NewSession()
	if err != nil {
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0, // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	err = c.Session.session.RequestPty("xterm", 80, 40, modes)
	if err != nil {
		return
	}
	return
}

func (c *SSHCli) Close() (err error) {
	err = c.closeSession()
	err = c.closeClient()
	return
}

func (c *SSHCli) closeSession() (err error) {
	if c.Session.session != nil {
		err = c.Session.session.Close()
		c.Session.session = nil
	}
	return
}
package bssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"

	"github.com/bypdhu/ssh-executor/task"
)

type SSHCli struct {
	Cli
	Session
	task.Task
}

type Session struct {
	session *ssh.Session
}

func New(ip string, port int, username string, password string) *SSHCli {
	cli := new(SSHCli)
	cli.Host = ip
	cli.Port = port
	cli.Username = username
	cli.Password = password

	return cli
}

func (c *SSHCli) SshRun() (err error) {
	if c.Session.session == nil || c.client == nil {
		if err = c.newSession(); err != nil {
			return
		}
	}

	defer c.closeSession()

	c.ExitCode = 0

	buf, err := c.Session.session.CombinedOutput(c.Command)
	if err != nil {
		if v, ok := err.(*ssh.ExitError); ok {
			c.ExitCode = v.Waitmsg.ExitStatus()
		}
		c.ExitCode = -1
	}
	c.Result = fmt.Sprintf("%s", buf)
	//c.Session.LastResult = string(buf)

	return
}

func (c *SSHCli) RunCommand(cmd string) (error) {
	c.OriginalCommand = cmd
	_cmd := ""
	if c.ShellArgs.Become {
		_cmd += c.ShellArgs.BecomeMethod + " -H -S -n "
		if c.ShellArgs.BecomeUser != "" {
			_cmd += "-u " + c.ShellArgs.BecomeUser + " "
		}
	}
	_cmd += "/bin/sh "
	if c.ShellArgs.Login {
		_cmd += " --login  "
	}
	_cmd += " -c '"
	_cmd += cmd + "'"
	c.Command = _cmd
	return c.SshRun()
}

func (c *SSHCli) RunCommand2(cmd, becomeMethod, becomeUser string, become, login bool) (error) {
	_cmd := ""
	if become {
		_cmd += becomeMethod + " -H -S -n "
		if becomeUser != "" {
			_cmd += "-u " + becomeMethod + " "
		}
	}
	_cmd += "/bin/sh "
	if login {
		_cmd += " --login  "
	}
	_cmd += " -c '"
	_cmd += cmd + "'"
	c.Command = _cmd
	return c.SshRun()
}

func (c *SSHCli) newSession() (err error) {
	if c.client == nil {
		if err = c.newClient(); err != nil {
			return
		}
	}

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
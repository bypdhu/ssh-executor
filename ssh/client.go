package bssh

import (
	"fmt"
	"net"
	"time"
	"golang.org/x/crypto/ssh"
)

type Cli struct {
	Host     string
	Port     int
	Username string
	Password string
	client   *ssh.Client
}

type SSHCli struct {
	Cli
	SSHConfig
	Session Session
}

type SFTPCli struct {
	Cli
	SSHConfig
}

type Session struct {
	session    *ssh.Session
	LastCmd    string
	LastResult string
	ExitCode   int
}

type SSHConfig struct {
	Timeout int
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
	c.Session.LastCmd = cmd
	buf, err := c.Session.session.CombinedOutput(cmd)
	if err != nil {
		if v, ok := err.(*ssh.ExitError); ok {
			c.Session.ExitCode = v.Waitmsg.ExitStatus()
		}
		c.Session.ExitCode = -1
	}
	c.Session.LastResult = string(buf)

	return
}

func (c *SSHCli) newClient() (err error) {
	var (
		auth []ssh.AuthMethod
		config ssh.Config
		clientConfig *ssh.ClientConfig
	)
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(c.Password))

	config = ssh.Config{
		Ciphers:[]string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128",
			"aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
	}
	clientConfig = &ssh.ClientConfig{
		User: c.Username,
		Auth: auth,
		Timeout: time.Duration(c.SSHConfig.Timeout),
		Config: config,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	c.client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), clientConfig)
	if err != nil {
		return
	}

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
	err = c.closeSession()
	return
}

func (c *SSHCli) closeClient() (err error) {
	if c.client != nil {
		err = c.client.Close()
		c.client = nil
	}
	return
}

func (c *SSHCli) closeSession() (err error) {
	if c.Session.session != nil {
		err = c.Session.session.Close()
		c.Session.session = nil
	}
	return
}
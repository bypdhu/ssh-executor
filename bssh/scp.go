package bssh

import (
	"os"
	"io"

	"github.com/pkg/sftp"
)

type SFTPCli struct {
	Cli
	SftpClient *sftp.Client
}

func NewSftp(ip string, port int, username string, password string) (*SFTPCli) {
	cli := new(SFTPCli)
	cli.Host = ip
	cli.Port = port
	cli.Username = username
	cli.Password = password
	return cli
}

func (c *SFTPCli) newSftpClient() (err error) {
	c.SftpClient, err = sftp.NewClient(c.client)
	if err != nil {
		c.SftpClient = nil
		return
	}
	return
}

func (c *SFTPCli) CopyOneFile(mode string, src string, dest string) (err error) {
	if c.client == nil {
		if err = c.newClient(); err != nil {
			return
		}
	}
	if c.SftpClient == nil {
		if err = c.newSftpClient(); err != nil {
			return
		}
	}
	switch mode {
	case "pull":
		err = c.PullFile(src, dest)
	case "push":
		err = c.PushFile(src, dest)
	}
	return
}

func (c *SFTPCli) PullFile(remote string, local string) (err error) {
	r, err := c.SftpClient.Open(remote)
	if err != nil {
		return
	}
	defer r.Close()

	l, err := os.Create(local)
	if err != nil {
		return
	}
	defer l.Close()

	if _, err = r.WriteTo(l); err != nil {
		return
	}

	return
}

func (c *SFTPCli) PushFile(local string, remote string) (err error) {
	l, err := os.Open(local)
	if err != nil {
		return
	}
	defer l.Close()

	r, err := c.SftpClient.Create(remote)
	if err != nil {
		return
	}
	defer r.Close()

	bufSize := 1024

	buf := make([]byte, bufSize)
	for {
		var n int
		n, err = l.Read(buf)
		//fmt.Printf("n:%d\n", n)
		//fmt.Printf("Error:%s\n", err)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return
		}
		if n == 0 {
			// perhaps no use.
			break
		}
		if n < bufSize {
			buf = buf[:n]
		}
		_, err = r.Write(buf)
		if err != nil {
			return
		}
	}
	return
}

func (c *SFTPCli) Close() (err error) {
	err = c.closeSftp()
	err = c.closeClient()
	return
}

func (c *SFTPCli) closeClient() (err error) {
	if c.client != nil {
		err = c.client.Close()
		c.client = nil
	}
	return
}

func (c *SFTPCli) closeSftp() (err error) {
	if c.SftpClient != nil {
		err = c.SftpClient.Close()
		c.SftpClient = nil
	}
	return
}
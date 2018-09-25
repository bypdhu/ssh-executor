package bssh

import (
	"os"
	"io"

	"github.com/pkg/sftp"
	"github.com/pkg/errors"
	"github.com/bypdhu/ssh-executor/utils"
)

type SFTPCli struct {
	SSHCli
	SftpClient *sftp.Client
	SFTPConfig
}

type SFTPConfig struct {
	Md5       string
	ForceCopy bool
	BufSize   int
}

type SrcDest struct {
	SrcDestResult
	Src  string
	Dest string
}

type SrcDestResult struct {
	Change bool
	err    error
}

type SftpMode string

const (
	SftpPull SftpMode = "PULL"
	SftpPush SftpMode = "PUSH"
)

func NewSftp(ip string, port int, username string, password string) (*SFTPCli) {
	cli := new(SFTPCli)
	cli.Host = ip
	cli.Port = port
	cli.Username = username
	cli.Password = password

	cli.ForceCopy = false
	cli.BufSize = 1024 * 1024

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

func (c *SFTPCli) CopyOneFile(mode SftpMode, src string, dest string) (change bool, err error) {
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
	case SftpPull:
		change, err = c.PullFile(src, dest)
	case SftpPush:
		change, err = c.PushFile(src, dest)
	}
	return
}

func (c *SFTPCli) CopyManyFiles(mode SftpMode, srcDests []*SrcDest) {
	for _, srcDest := range srcDests {
		srcDest.Change, srcDest.err = c.CopyOneFile(mode, srcDest.Src, srcDest.Dest)
	}
}

func (c *SFTPCli) PullFile(remote string, local string) (bool, error) {

	if utils.IsFile(local) && c.ForceCopy == false {
		l_md5, _ := utils.GetMd5FromPath(local)
		r_md5, err := getRemoteFileMd5(c, remote)
		if err != nil {
			return false, errors.New(err.Error() + ". Detail: " + c.Session.LastResult)
		}
		if r_md5 != "" && r_md5 == l_md5 {
			return false, nil
		}
	}

	r, err := c.SftpClient.Open(remote)
	if err != nil {
		return false, err
	}
	defer r.Close()

	l, err := os.Create(local)
	if err != nil {
		return false, err
	}
	defer l.Close()

	if _, err = r.WriteTo(l); err != nil {
		return false, err
	}

	return true, nil
}

func (c *SFTPCli) PushFile(local string, remote string) (bool, error) {
	l, err := os.Open(local)
	if err != nil {
		return false, err
	}
	defer l.Close()

	if c.ForceCopy == false {
		l_md5, err := utils.GetMd5FromPath(local)
		if err != nil {
			return false, err
		}
		r_md5, _ := getRemoteFileMd5(c, remote)

		if l_md5 != "" && r_md5 == l_md5 {
			return false, nil
		}
	}

	r, err := c.SftpClient.Create(remote)
	if err != nil {
		return false, err
	}
	defer r.Close()

	buf := make([]byte, c.BufSize)
	for {
		var n int
		n, err = l.Read(buf)
		//fmt.Printf("n:%d\n", n)
		//fmt.Printf("Error:%s\n", err)
		if err != nil {
			if err == io.EOF {
				return true, nil
			}
			return false, err
		}
		if n == 0 {
			// perhaps no use.
			return true, nil
		}
		if n < c.BufSize {
			buf = buf[:n]
		}
		_, err = r.Write(buf)
		if err != nil {
			return false, err
		}
	}
}

func (c *SFTPCli) Close() (err error) {
	err = c.closeSftp()
	err = c.closeClient()
	return
}

func (c *SFTPCli) closeSftp() (err error) {
	if c.SftpClient != nil {
		err = c.SftpClient.Close()
		c.SftpClient = nil
	}
	return
}


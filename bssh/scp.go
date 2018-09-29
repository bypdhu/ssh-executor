package bssh

import (
	"os"
	"io"
	"fmt"

	"github.com/pkg/sftp"
	"github.com/pkg/errors"
	"github.com/bypdhu/ssh-executor/utils"
	"github.com/bypdhu/ssh-executor/common"
	"github.com/bypdhu/ssh-executor/task"
	"strings"
)

// Sftp client
type SFTPCli struct {
	SSHCli
	sftpClient *sftp.Client
	SFTPConfig
}

type SFTPConfig struct {
	BufSize int
}

func NewSftp(ip string, port int, username string, password string) (c *SFTPCli) {
	c = new(SFTPCli)
	c.Host = ip
	c.Port = port
	c.Username = username
	c.Password = password

	// SFTPConfig
	c.BufSize = 1024 * 1024

	if err := c.newSftpClient(); err != nil {
		return
	}

	return
}

func (c *SFTPCli) newSftpClient() (err error) {
	if c.client == nil {
		if err = c.newClient(); err != nil {
			return
		}
	}
	c.sftpClient, err = sftp.NewClient(c.client)
	if err != nil {
		c.sftpClient = nil
		return
	}
	return
}

func (c *SFTPCli) SftpRun() error {
	return c.CopyManyFiles()
}

func (c *SFTPCli) CopyManyFiles() error {
	mode := c.SftpMode
	errsEntity := []*task.CopyOneFile{}

	for _, srcDest := range c.CopyFiles {
		err := c.CopyOneFile(mode, srcDest)
		if err != nil {
			if c.CopyArgs.IgnoreErr {
				errsEntity = append(errsEntity, srcDest)
				continue
			}
			return err
		}
	}
	if len(errsEntity) != 0 {
		return getErrorFromSrcDest(errsEntity)
	}
	return nil
}

func getErrorFromSrcDest(sds []*task.CopyOneFile) (error) {
	s := ""
	for _, sd := range sds {
		if sd.Err == nil {
			continue
		}
		s += fmt.Sprintf("Copy %s to %s, err %s. \n ", sd.Src, sd.Dest, sd.Err)
	}
	return errors.New(s)
}

func (c *SFTPCli) CopyOneFile(mode common.SftpMode, one *task.CopyOneFile) (err error) {
	if c.sftpClient == nil || c.client == nil {
		if err = c.newSftpClient(); err != nil {
			return
		}
	}

	switch mode {
	case common.SFTP_PULL:
		one.Changed, one.Err = c.PullFile(one)
	case common.SFTP_PUSH:
		one.Changed, one.Err = c.PushFile(one)
	}
	err = one.Err
	return
}

func (c *SFTPCli) PullFile(one *task.CopyOneFile) (bool, error) {
	local := one.Dest
	remote := one.Src

	if utils.IsDir(local) {
		return false, errors.New("local: " + local + " is a dir.")
	}

	if utils.IsFile(local) && one.ForceCopy == false {
		l_md5, _ := utils.GetMd5FromPath(local)
		r_md5, err := c.GetRemoteFileMd5(remote)
		if err != nil {
			return false, errors.New(err.Error() + ". Detail: " + c.Result)
		}
		if r_md5 != "" && r_md5 == l_md5 {
			return false, nil
		}
	}

	r, err := c.sftpClient.Open(remote)
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

func (c *SFTPCli) PushFile(one *task.CopyOneFile) (bool, error) {
	local := one.Src
	remote := one.Dest

	if utils.IsDir(local) {
		return false, errors.New("local: " + local + " is a dir.")
	}

	//if one.CreateDirectory {
	//	base, toCreate := c.GetDirExists(remote)
	//
	//	fmt.Printf("base:%s,toCreate:%s\n", base, toCreate)
	//	if err := c.CreateDirsRemote(one, base, toCreate); err != nil {
	//		return false, errors.New(err.Error() + ". Detail: " + c.Result)
	//	}
	//	if err := c.ChmodRemote(one, base, toCreate); err != nil {
	//		return false, errors.New(err.Error() + ". Detail: " + c.Result)
	//	}
	//	if err := c.ChownRemote(one, base, toCreate); err != nil {
	//		return false, errors.New(err.Error() + ". Detail: " + c.Result)
	//	}
	//}

	l, err := os.Open(local)
	if err != nil {
		return false, err
	}
	defer l.Close()

	if one.ForceCopy == false {
		l_md5, err := utils.GetMd5FromPath(local)
		if err != nil {
			return false, err
		}
		r_md5, _ := c.GetRemoteFileMd5(remote)

		if l_md5 != "" && r_md5 == l_md5 {
			return false, nil
		}
	}

	r, err := c.sftpClient.Create(remote)
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
	if c.sftpClient != nil {
		err = c.sftpClient.Close()
		c.sftpClient = nil
	}
	return
}

func (c *SFTPCli) GetDirExists(path string) (base string, toCreate string) {
	path, _ = sftp.Split(path)
	ps := strings.Split(path, "/")
	//fmt.Println(ps)
	for i := range ps {
		base = strings.Join(ps[:i + 1], "/")
		toCreate = strings.Join(ps[i + 1:], "/")
		//fmt.Printf("base:%s, create:%s\n", base, toCreate)
		if base == "" {
			continue
		}
		_, err := c.sftpClient.ReadDir(base)
		if err != nil {
			base = strings.Join(ps[:i ], "/")
			toCreate = strings.Join(ps[i:], "/")
			return
		}
	}
	return
}

func (c *SFTPCli) CreateDirsRemote(one *task.CopyOneFile, cd string, path string) (err error) {
	if path == "" {
		return
	}

	s := []string{}
	if one.Become {
		s = append(s, one.BecomeMethod)
	}
	s = append(s,"/bin/sh -c")
	if cd != "" {
		s = append(s, "'cd " + cd + " &&")
	} else {
		s = append(s, "'")
	}
	s = append(s, "mkdir")
	s = append(s, path)
	s = append(s, "'")

	fmt.Println(strings.Join(s, " "))
	err = c.SSHCli.RunCommand(strings.Join(s, " "))
	one.Result = c.Result
	return

}

func (c *SFTPCli) ChmodRemote(one *task.CopyOneFile, cd string, path string) (err error) {
	if path == "" {
		return
	}

	s := []string{}
	if one.Become {
		s = append(s, one.BecomeMethod)
	}
	s = append(s,"/bin/sh -c")
	if cd != "" {
		s = append(s, "'cd " + cd + " &&")
	} else {
		s = append(s, "'")
	}
	s = append(s, "chmod -R")
	s = append(s, one.DirectoryMode)
	s = append(s, path)
	s = append(s, "'")

	fmt.Println(strings.Join(s, " "))
	err = c.SSHCli.RunCommand(strings.Join(s, " "))
	return
}

func (c *SFTPCli) ChownRemote(one *task.CopyOneFile, cd string, path string) (err error) {
	if path == "" {
		return
	}

	s := []string{}
	if one.Become {
		s = append(s, one.BecomeMethod)
	}
	s = append(s,"/bin/sh -c")
	if cd != "" {
		s = append(s, "'cd " + cd + " &&")
	} else {
		s = append(s, "'")
	}
	s = append(s, "chown -R")
	s = append(s, one.Owner + ":" + one.Group)
	s = append(s, path)
	s = append(s, "'")

	fmt.Println(strings.Join(s, " "))
	err = c.SSHCli.RunCommand(strings.Join(s, " "))
	return
}

package bssh

import (
	"strings"
	"os"
)

func (c *SSHCli) GetRemoteFileMd5(remote string) (s string, err error) {
	err = c.RunCommandDirect("md5sum " + remote)
	if err != nil {
		return
	}

	r := c.Stdout
	s = strings.Split(r, " ")[0]
	return
}

func (c *SSHCli) ChmodDirRemote(cd string, path string, mode os.FileMode) (err error) {
	return
}


//func (c *SFTPCli) ChmodRemote(path string, mode os.FileMode) (err error) {
//	err = c.sftpClient.Chmod(path, mode)
//	return
//}

//func (c *SFTPCli) ChownRemote(cd string, path string, user string, group string) (err error) {
//
//	err = c.RunCommand("id -u " + user)
//	if err != nil {
//		return
//	}
//	r := c.Result // like  601\r\n
//	uid, err := strconv.Atoi(strings.TrimSpace(r))
//	if err != nil {
//		return
//	}
//
//	err = c.RunCommand("id -g " + user)
//	if err != nil {
//		return
//	}
//	r = c.Result
//	gid, err := strconv.Atoi(strings.TrimSpace(r))
//	if err != nil {
//		return
//	}
//
//	err = c.sftpClient.Chown(path, uid, gid)
//
//	return
//}
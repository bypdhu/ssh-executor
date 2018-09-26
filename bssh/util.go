package bssh

import (
	"strings"
)

func getRemoteFileMd5(c *SFTPCli, remote string) (string, error) {
	err := c.SSHCli.Run("md5sum " + remote)
	if err != nil {
		return "", err
	}

	r := c.Session.Result
	rl := strings.Split(r, " ")
	return rl[0], nil
}
package conf

import (
    "strings"

    "git.eju-inc.com/ops/go-common/log"

    "github.com/bypdhu/ssh-executor/utils"
)

func GetHostsFromConfig(c *Config) ([]string) {
    hosts := strings.Split(c.Direct.Hosts, ",")
    if c.Direct.HostsFile != "" {
        hosts_from_file, err := utils.ReadLineNotEmpty(c.Direct.HostsFile)
        if err != nil {
            log.Error(err)
        }
        hosts = append(hosts, hosts_from_file...)
    }

    hosts = utils.RemoveDupString(hosts)

    return hosts
}

func CopySSHConfig(from, to *SSHConfig) {
    if from.Timeout != 0 {
        to.Timeout = from.Timeout
    }
    to.UserName = from.UserName
    to.Password = from.Password
}
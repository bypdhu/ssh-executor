package conf

import (
	"strings"
	"github.com/bypdhu/ssh-executor/utils"
	"git.eju-inc.com/ops/go-common/log"
)

func GetHosts(c *Config) ([]string) {
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
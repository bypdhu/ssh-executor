package main

import (
	"os"
	"strings"
	"fmt"

	"git.eju-inc.com/ops/go-common/log"
	"git.eju-inc.com/ops/go-common/log/ctx"
	"git.eju-inc.com/ops/go-common/version"

	"github.com/bypdhu/ssh-executor/cmd/ssh-executor/flags"
	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/utils"
	"github.com/bypdhu/ssh-executor/result"
	"github.com/bypdhu/ssh-executor/module"
)

var (
	c *conf.Config
	results map[string]*result.BaseResult
)

func main() {
	fs := flags.ParseFlags(os.Args)
	log.SetContextExtractor(ctx.ZipkinTraceExtractor{})

	c = conf.Load(fs.ConfigFilePath)
	//log.Infof("config from file and default: %s", c)

	flags.OverrideConfWithFlags(c, fs)
	//log.Infof("final config, override by command-args: %s", c)
	//log.Infof("==========launch type: %s\n", c.LaunchType)

	switch c.LaunchType {
	case "direct":
		doDirect(c)
	case "server":
		log.Infof("Starting up %s ...\n", version.Print())
		log.Infoln(fs.PrintAll())
		log.Infof("Final config, override by command-args: %s", c)

		doServer(c)
	}
}

func doServer(c *conf.Config) {

}

func doDirect(c *conf.Config) {
	results = make(map[string]*result.BaseResult)

	hosts := getHosts(c)

	module.RunAll(c, hosts, results)

	for _, host := range hosts {
		fmt.Printf("host:%s, result:%s, exitCode:%d, err:%s\n", host, results[host].Result, results[host].ExitCode, results[host].Err)
	}
}

func getHosts(c *conf.Config) ([]string) {
	hosts := strings.Split(c.Direct.Hosts, ",")
	if c.Direct.HostsFile != "" {
		hosts_from_file, err := utils.ReadLineNotEmpty(c.Direct.HostsFile)
		if err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, hosts_from_file...)
	}

	hosts = utils.RemoveDupString(hosts)

	return hosts
}
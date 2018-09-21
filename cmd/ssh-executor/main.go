package main

import (
	"os"
	"sync"
	"strings"
	"fmt"

	"git.eju-inc.com/ops/go-common/log"
	"git.eju-inc.com/ops/go-common/log/ctx"
	"git.eju-inc.com/ops/go-common/version"

	"github.com/bypdhu/ssh-executor/cmd/ssh-executor/flags"
	"github.com/bypdhu/ssh-executor/conf"
	"github.com/bypdhu/ssh-executor/ssh"
	"github.com/bypdhu/ssh-executor/utils"
)

type Result struct {
	result   string
	exitCode int
	err      error
}

var (
	c *conf.Config
	results map[string]*Result
	wg sync.WaitGroup
)

func main() {
	fs := flags.ParseFlags(os.Args)
	log.SetContextExtractor(ctx.ZipkinTraceExtractor{})
	//log.Infof("Starting up %s ...\n", version.Print())
	//log.Infoln(fs.PrintAll())

	c = conf.Load(fs.ConfigFilePath)
	//log.Infof("config from file and default: %s", c)

	overrideConfWithFlags(c, fs)

	//log.Infof("final config, override by command-args: %s", c)
	//log.Infof("==========launch type: \n", c.LaunchType)

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
	results = make(map[string]*Result)
	hosts := strings.Split(c.Direct.Hosts, ",")
	if c.Direct.HostsFile != "" {
		hosts_from_file, err := utils.ReadLineNotEmpty(c.Direct.HostsFile)
		if err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, hosts_from_file...)
	}

	hosts = utils.RemoveDupString(hosts)

	for _, host := range hosts {
		results[host] = &Result{}
		wg.Add(1)
		go runCmd(c, host, results[host])
	}

	wg.Wait()

	for _, host := range hosts {
		fmt.Printf("host:%s, result:%s, exitCode:%s, err:%s\n", host, results[host].result, results[host].exitCode, results[host].err)
	}
}

func runCmd(c *conf.Config, h string, result *Result) {
	defer wg.Done()
	client := bssh.New(h, 22, c.Direct.UserName, c.Direct.Password)
	//log.Infof("+++++++++now run %s on host %s\n", c.Direct.Command, h)
	//log.Infof("client is %s", client)
	result.err = client.Run(c.Direct.Command)
	result.result = client.Session.LastResult
	result.exitCode = client.Session.ExitCode

	//log.Infof("++++++result is %s on host %s\n", result.result, h)
}

func overrideConfWithFlags(c *conf.Config, i flags.CommandLineFlags) {
	c.LaunchType = i.LaunchType
	c.SSHConfig.Timeout = i.SshTimeout

	switch i.LaunchType {
	case "direct":
		if i.Hosts != "" {
			c.Direct.Hosts = i.Hosts
		}
		if i.HostsFile != "" {
			c.Direct.HostsFile = i.HostsFile
		}
		if i.UserName != "" {
			c.Direct.UserName = i.UserName
		}
		if i.Password != "" {
			c.Direct.Password = i.Password
		}
		if i.Command != "" {
			c.Direct.Command = i.Command
		}
	case "server":
		if i.WebAddress != "" {
			c.Serv.Web.ListenAddress = i.WebAddress
		}
		if i.TelemetryAddress != "" {
			c.Serv.Telemetry.ListenAddress = i.TelemetryAddress
		}
	}

}
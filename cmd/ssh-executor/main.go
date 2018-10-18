package main

import (
    "os"
    "strings"

    "git.eju-inc.com/ops/go-common/log"
    "git.eju-inc.com/ops/go-common/log/ctx"
    "git.eju-inc.com/ops/go-common/version"

    "github.com/bypdhu/ssh-executor/cmd/ssh-executor/flags"
    "github.com/bypdhu/ssh-executor/conf"
    "github.com/bypdhu/ssh-executor/common"
    "github.com/bypdhu/ssh-executor/direct"
    "github.com/bypdhu/ssh-executor/web"
)

var (
    c *conf.Config
    //cs map[string]*conf.Config
    //rs map[string]*result.DirectResult
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
    case common.LAUNCH_DIRECT.String(), strings.ToLower(common.LAUNCH_DIRECT.String()):
        direct.Run(c)

    case common.LAUNCH_SERVER.String(), strings.ToLower(common.LAUNCH_SERVER.String()):
        log.Infof("Starting up %s ...\n", version.Print())
        log.Infoln(fs.PrintAll())
        log.Infof("Final config, override by command-args: %s", c)

        web.Run(c)
    }
}


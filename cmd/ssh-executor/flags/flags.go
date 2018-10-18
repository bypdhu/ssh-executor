package flags

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"

    kingpin "gopkg.in/alecthomas/kingpin.v2"
    "github.com/pkg/errors"

    "git.eju-inc.com/ops/go-common/version"
    "git.eju-inc.com/ops/go-common/log"

    "github.com/bypdhu/ssh-executor/common"
    "github.com/bypdhu/ssh-executor/conf"
)

type CommandLineFlags struct {
    ConfigFilePath   string
    LaunchType       string

    SSHTimeout       int
    UserName         string
    Password         string

    WebAddress       string // when LaunchType=server
    TelemetryAddress string // when LaunchType=server

    Hosts            string // when LaunchType=direct. separate by ","
    HostsFile        string // when LaunchType=direct. One Ip on line.
    Module           string // when LaunchType=direct.
    Command          string // when LaunchType=direct.
}

func (f *CommandLineFlags) PrintAll() (content string) {
    switch f.LaunchType {
    case common.LAUNCH_SERVER.String(), strings.ToLower(common.LAUNCH_SERVER.String()):
        content = fmt.Sprintf(
            `Flags{
            launch.type=%s,
            config.file=%s,
            web.listen_address=%s,
            telemetry.listen_address=%s,
            ssh.timeout=%d,
            }
            `,
            f.LaunchType,
            f.ConfigFilePath,
            f.WebAddress,
            f.TelemetryAddress,
            f.SSHTimeout,
        )
    case common.LAUNCH_DIRECT.String(), strings.ToLower(common.LAUNCH_DIRECT.String()):
        content = fmt.Sprintf(
            `Flags{
            launch.type=%s,
            config.file=%s,
            hosts=%s,
            hosts.file=%s,
            user.name=%s,
            user.pass=***,
            module=%s,
            command=%s,
            ssh.timeout=%d,
            }
            `,
            f.LaunchType,
            f.ConfigFilePath,
            f.Hosts,
            f.HostsFile,
            f.UserName,
            f.Module,
            f.Command,
            f.SSHTimeout,
        )
    default:
        content = fmt.Sprintf(
            `Flags{
            launch.type=%s,
            config.file=%s,
            ssh.timeout=%d,
            } required launch.type=server or launch.type=direct.
            `,
            f.LaunchType,
            f.ConfigFilePath,
            f.SSHTimeout,
        )
    }
    return
}

func ParseFlags(args []string) (clfs CommandLineFlags) {
    a := kingpin.New(filepath.Base(args[0]), version.AppName)

    a.Version(version.Print())
    a.HelpFlag.Short('h')
    log.AddFlags(a)

    a.Flag("config.file", "application's configuration file path.").
            Default("").Short('c').StringVar(&clfs.ConfigFilePath)
    a.Flag("launch.type", "server/direct;default direct. server will setup a http server. direct will execute command once.").
            Default(common.LAUNCH_DIRECT.String()).Short('T').StringVar(&clfs.LaunchType)
    a.Flag("ssh.timeout", "timeout in ssh connection. default 30s.").
            Default("30").Short('t').IntVar(&clfs.SSHTimeout)
    a.Flag("web.listen_address", "[launch.type=server] Address to listen on for UI, API.").
            Default("").Short('a').StringVar(&clfs.WebAddress)
    a.Flag("telemetry.listen_address", "[launch.type=server] Address to listen on for telemetry.").
            Default("").StringVar(&clfs.TelemetryAddress)
    a.Flag("hosts", "[launch.type=direct] Hosts to connect by ssh. Combined by ','. Add hosts.file.").
            Default("").Short('i').StringVar(&clfs.Hosts)
    a.Flag("hosts.file", "[launch.type=direct] File of hosts to connect by ssh. One ip on line. Add hosts.").
            Default("").Short('f').StringVar(&clfs.HostsFile)
    a.Flag("user.name", "[launch.type=direct] Username for ssh connection.").
            Default("").Short('u').StringVar(&clfs.UserName)
    a.Flag("user.pass", "[launch.type=direct] Password for ssh connection.").
            Default("").Short('p').StringVar(&clfs.Password)
    a.Flag("module", "[launch.type=direct] Module to handle. like 'shell' 'copy' ").
            Default(common.MODULE_SHELL.String()).Short('m').StringVar(&clfs.Module)
    a.Flag("command", "[launch.type=direct] Command to handle.").
            Default("").Short('C').StringVar(&clfs.Command)

    _, err := a.Parse(args[1:])
    if err != nil {
        a.Usage(args[1:])
        fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
        os.Exit(2)
    }
    return
}

func OverrideConfWithFlags(c *conf.Config, i CommandLineFlags) {
    c.LaunchType = i.LaunchType

    if i.SSHTimeout != 0 {
        c.SSHConfig.Timeout = i.SSHTimeout
    }
    if i.UserName != "" {
        c.SSHConfig.UserName = i.UserName
    }
    if i.Password != "" {
        c.SSHConfig.Password = i.Password
    }

    switch i.LaunchType {
    case common.LAUNCH_DIRECT.String(), strings.ToLower(common.LAUNCH_DIRECT.String()):
        if i.Hosts != "" {
            c.Direct.Hosts = i.Hosts
        }
        if i.HostsFile != "" {
            c.Direct.HostsFile = i.HostsFile
        }
        if i.Module != "" {
            c.Direct.Module = i.Module
        }
        if i.Command != "" {
            c.Direct.Command = i.Command
        }
    case common.LAUNCH_SERVER.String(), strings.ToLower(common.LAUNCH_SERVER.String()):
        if i.WebAddress != "" {
            c.Serv.Web.ListenAddress = i.WebAddress
        }
        if i.TelemetryAddress != "" {
            c.Serv.Telemetry.ListenAddress = i.TelemetryAddress
        }
    }

}
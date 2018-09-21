package flags

import (
	"fmt"
	"os"
	"path/filepath"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/pkg/errors"
	"git.eju-inc.com/ops/go-common/version"
	"git.eju-inc.com/ops/go-common/log"
)

type CommandLineFlags struct {
	ConfigFilePath   string
	LaunchType       string
	SshTimeout       int

	WebAddress       string // when LaunchType=server
	TelemetryAddress string // when LaunchType=server

	Hosts            string // when LaunchType=direct. separate by ","
	HostsFile        string // when LaunchType=direct. One Ip on line.
	UserName         string // when LaunchType=direct.
	Password         string // when LaunchType=direct.
	Command          string // when LaunchType=direct.
}

func (f *CommandLineFlags) PrintAll() (content string) {
	switch f.LaunchType {
	case "server":
		content = fmt.Sprintf(
			`Flags{
			config.file=%s,
			launch.type=%s,
			web.listen_address=%s,
			telemetry.listen_address=%s,
			ssh.timeout=%d,
			}
			`,
			f.ConfigFilePath,
			f.LaunchType,
			f.WebAddress,
			f.TelemetryAddress,
			f.SshTimeout,
		)
	case "direct":
		content = fmt.Sprintf(
			`Flags{
			config.file=%s,
			launch.type=%s,
			hosts=%s,
			hosts.file=%s,
			user.name=%s,
			user.pass=***,
			ssh.timeout=%d,
			}
			`,
			f.ConfigFilePath,
			f.LaunchType,
			f.Hosts,
			f.HostsFile,
			f.UserName,
			f.SshTimeout,
		)
	default:
		content = fmt.Sprintf(
			`Flags{
			config.file=%s,
			launch.type=%s,
			ssh.timeout=%d,
			} required launch.type=server or launch.type=direct.
			`,
			f.ConfigFilePath,
			f.LaunchType,
			f.SshTimeout,
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
			Default("direct").Short('T').StringVar(&clfs.LaunchType)
	a.Flag("ssh.timeout", "timeout in ssh connection. default 30s.").
			Default("30").Short('t').IntVar(&clfs.SshTimeout)
	a.Flag("web.listen_address", "[launch.type=server] Address to listen on for UI, API.").
			Default("").StringVar(&clfs.WebAddress)
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
	a.Flag("command", "[launch.type=direct] Command for ssh connection.").
			Default("").Short('C').StringVar(&clfs.Command)

	_, err := a.Parse(args[1:])
	if err != nil {
		a.Usage(args[1:])
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
		os.Exit(2)
	}
	return
}
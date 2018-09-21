package conf

import (
	"fmt"
	"sync"

	"git.eju-inc.com/ops/go-common/config"
)

type Config struct {
	Serv       Server    `yaml:"server"`
	Direct     Direct    `yaml:"direct"`
	LaunchType string    `yaml:"launch_type"`
	SSHConfig  SSHConfig `yaml:"ssh_config"`

	original   string
}

type SSHConfig struct {
	Timeout int `yaml:"timeout"`
}

type Server struct {
	Web       Web       `yaml:"web"`       // when LaunchType=server
	Telemetry Telemetry `yaml:"telemetry"` // when LaunchType=server
}

type Web struct {
	ListenAddress string `yaml:"listen_address"`
}
type Telemetry struct {
	ListenAddress string `yaml:"listen_address"`
}

type Direct struct {
	Hosts     string `yaml:"hosts"`      // when LaunchType=direct. separate by ","
	HostsFile string `yaml:"hosts_file"` // when LaunchType=direct. separate by ","
	UserName  string `yaml:"username"`   // when LaunchType=direct.
	Password  string `yaml:"password"`   // when LaunchType=direct.
	Command   string `yaml:"command"`    // when LaunchType=direct.
}

var (
	DefaultConfig = Config{
		Serv:Server{
			Web:Web{ListenAddress:"localhost:9888", },
			Telemetry:Telemetry{ListenAddress:"localhost:9889", },
		},
		Direct:Direct{},
		LaunchType:"direct",
	}
	cfg *Config
	mtx sync.RWMutex
)

func Load(configFile string) *Config {
	mtx.Lock()
	defer mtx.Unlock()

	cfg = &DefaultConfig
	if configFile != "" {
		cfg.original = config.LoadConfig(configFile, cfg)
	}
	return cfg
}

func Get() *Config {
	mtx.RLock()
	defer mtx.RUnlock()

	return cfg
}

func (w Web) String() string {
	return fmt.Sprintf("{ListenAddress:%s}", w.ListenAddress)
}
func (t Telemetry) String() string {
	return fmt.Sprintf("{ListenAddress:%s}", t.ListenAddress)
}
func (s Server) String() string {
	return fmt.Sprintf("{Web:%s,Telemetry:%s}", s.Web, s.Telemetry)
}
func (d Direct) String() string {
	return fmt.Sprintf("{Hosts:%s,HostsFile:%s,UserName:%s,Password:%s,Command:%s}", d.Hosts, d.HostsFile, d.UserName, d.Password, d.Command)
}
func (d Direct) StringSecure() string {
	return fmt.Sprintf("{Hosts:%s,HostsFile:%s,UserName:%s,Password:***,Command:%s}", d.Hosts, d.HostsFile, d.UserName, d.Command)
}
func (s SSHConfig) String() string {
	return fmt.Sprintf("{Timeout:%s}", s.Timeout)
}
func (c Config) String() string {
	return fmt.Sprintf("{LaunchType:%s,Server:%s,Direct:%s,SSHConfig:%s}", c.LaunchType, c.Serv, c.Direct, c.SSHConfig)
}
package conf

import (
	"fmt"
	"sync"
	"bytes"
	"encoding/json"

	"git.eju-inc.com/ops/go-common/config"

	"github.com/bypdhu/ssh-executor/common"
	"github.com/bypdhu/ssh-executor/task"
)

type Config struct {
	LaunchType string        `yaml:"launch_type"`

	Serv       Server        `yaml:"server"`

	Direct     Direct        `yaml:"direct"`

	SSHConfig  SSHConfig     `yaml:"ssh_config"`

	Tasks      []*task.Task  `yaml:"tasks"`

	original   string

	HostDup    string
}

type SSHConfig struct {
	Timeout  int    `yaml:"timeout" json:"timeout"`
	UserName string `yaml:"username" json:"username"`
	Password string `yaml:"password" json:"password"`
}

type Server struct {
	Web       Web       `yaml:"web"`
	Telemetry Telemetry `yaml:"telemetry"`
}

type Web struct {
	ListenAddress string `yaml:"listen_address"`
}

type Telemetry struct {
	ListenAddress string `yaml:"listen_address"`
}

type Direct struct {
	Hosts     string `yaml:"hosts"` // separate by ","
	HostsFile string `yaml:"hosts_file"`
	Module    string `yaml:"module"`
	Command   string `yaml:"command"`
}

var (
	DefaultConfig = Config{
		Serv:Server{
			Web:Web{ListenAddress:"localhost:9888", },
			Telemetry:Telemetry{ListenAddress:"localhost:9889", },
		},
		Direct:Direct{},
		LaunchType:common.LAUNCH_DIRECT.String(),
	}
	cfg *Config
	mtx sync.RWMutex
)

func GetCopiedConfigMap(c *Config, hs []string) map[string]*Config {
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(c)
	cs := make(map[string]*Config, len(hs))

	for _, h := range hs {
		_c := &Config{}
		json.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(_c)
		_c.HostDup = h
		cs[h] = _c
	}
	return cs
}

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
	return fmt.Sprintf("{Hosts:%s,HostsFile:%s,Command:%s}", d.Hosts, d.HostsFile, d.Command)
}
func (d Direct) StringSecure() string {
	return fmt.Sprintf("{Hosts:%s,HostsFile:%s,Command:%s}", d.Hosts, d.HostsFile, d.Command)
}
func (s SSHConfig) String() string {
	return fmt.Sprintf("{Timeout:%s,UserName:%s,Password:%s}", s.Timeout, s.UserName, s.Password)
}
func (s SSHConfig) StringSecure() string {
	return fmt.Sprintf("{Timeout:%s,UserName:%s,Password:***}", s.Timeout, s.UserName)
}
func (c Config) String() string {
	return fmt.Sprintf("{LaunchType:%s,Server:%s,Direct:%s,SSHConfig:%s}", c.LaunchType, c.Serv, c.Direct, c.SSHConfig)
}
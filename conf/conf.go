package conf

import (
	"github.com/BurntSushi/toml"
}

type LogConfig struct {
	Enable    bool   `toml:"enable"`
	Timestamp bool   `toml:"timestamp"`
	Dir       string `toml:"dirpath"`
}
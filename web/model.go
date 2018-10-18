package web

import (
    "github.com/bypdhu/ssh-executor/task"
    "github.com/bypdhu/ssh-executor/conf"
)

type WebBody struct {
    UserFlag  string          `json:"user_flag"`
    Hosts     []string        `json:"hosts"`
    SSHConfig conf.SSHConfig  `json:"ssh_config"`
    Tasks     []*task.Task    `json:"tasks"`
}


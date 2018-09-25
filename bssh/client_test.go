package bssh

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getUser() (username string, password string) {
	var _user User
	f, _ := ioutil.ReadFile("../tmp/config.yml")

	err := yaml.Unmarshal(f, &_user)
	if err != nil {
		fmt.Println(err)
	}
	username = _user.Username
	password = _user.Password
	return
}


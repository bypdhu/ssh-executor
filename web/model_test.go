package web

import (
	"testing"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"gopkg.in/yaml.v2"

	"github.com/bypdhu/ssh-executor/common"
	"github.com/bypdhu/ssh-executor/bssh"
)

type user struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func getUser() (username string, password string) {
	var _user user
	f, _ := ioutil.ReadFile("../tmp/config.yml")

	err := yaml.Unmarshal(f, &_user)
	if err != nil {
		fmt.Println(err)
	}
	username = _user.Username
	password = _user.Password
	return
}

func TestModel_WebBody(t *testing.T) {

	wb, _ := getWebBody()
	fmt.Printf("%+v\n", wb)

	for _, t := range wb.Tasks {
		fmt.Printf("%+v\n", t)
		switch t.Module {
		case common.MODULE_SHELL.String(), strings.ToLower(common.MODULE_SHELL.String()):
		case common.MODULE_COPY.String(), strings.ToLower(common.MODULE_COPY.String()):
			for _, f := range t.CopyFiles {
				fmt.Printf("%+v\n", f)
			}
		}
	}
}

func TestRun_TaskShell(t *testing.T) {
	wb, _ := getWebBody()
	username, password := getUser()

	client := bssh.New("10.99.70.38", 22, username, password)

	client.Task = wb.Tasks[0]
	client.RunCommandTask()

	fmt.Printf("%+v\n", client.Task)

}

func TestRun_TaskCopy(t *testing.T) {
	wb, _ := getWebBody()
	username, password := getUser()

	client := bssh.NewSftp("10.99.70.38", 22, username, password)

	client.Task = wb.Tasks[1]
	client.SftpStart()

	for _, i := range client.Task.CopyFiles {

		fmt.Printf("%+v\n", i)
	}
}

func getWebBody() (wb *WebBody, err error) {
	js := `{
  "user_flag": "general",
  "hosts": [
    "10.99.70.38",
    "10.99.70.35"
  ],
  "ssh_config": {
    "timeout": 32,
    "sh": "/bin/sh"
  },
  "tasks": [
    {
      "name": "ddd",
      "module": "shell",
      "args": {
        "shell": {
          "command": "ls /",
          "chdir": "/tmp",
          "login": true,
          "become": true,
          "become_user": "admin",
          "become_method": "sudo"
        }
      }
    },
    {
      "name": "dfafa",
      "module": "copy",
      "args": {
        "copy": {
          "ignore_err": true,
          "become": true,
          "become_user": "admin",
          "become_method": "sudo",
          "sftp_mode": "push",
          "copy_files": [
            {
              "src": "../tmp/config.yml",
              "dest": "/tmp/bian/config.model_test.yml",
              "owner": "admin",
              "group": "admin",
              "mode": "0644",
              "force": true,
              "create_directory": true,
              "recursive": true,
              "directory_mode": "0755"
            },
            {
              "src": "/tmp/foo2.yml",
              "dest": "foo2.yml",
              "owner": "foo2",
              "group": "foo2",
              "mode": "0644",
              "md5": "2414141"
            }
          ]
        }
      }
    }
  ]
}`
	wb = &WebBody{}
	err = json.Unmarshal([]byte(js), wb)
	if err != nil {
		fmt.Printf("err:%s\n", err)
	}
	return
}

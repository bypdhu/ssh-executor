package bssh

import (
	"testing"
	"fmt"
	"io/ioutil"
	"sync"
	"gopkg.in/yaml.v2"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func TestConnect(t *testing.T) {
	var user User
	f, _ := ioutil.ReadFile("../tmp/config.yml")

	err := yaml.Unmarshal(f, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("username:%s, password:%s\n", user.Username, user.Password)

	client := New("10.99.70.38", 22, user.Username, user.Password)

	for _, cmd := range []string{"pwd", "sss", "ps -ef|grep ansible", "ls"} {
		run(client, cmd)
	}
}

func TestSSHCli_Run(t *testing.T) {
	var (
		user User
		wg sync.WaitGroup
	)
	f, _ := ioutil.ReadFile("../tmp/config.yml")

	err := yaml.Unmarshal(f, &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("username:%s, password:%s\n", user.Username, user.Password)

	hosts := []string{"10.99.70.38", "10.99.70.34", "10.99.70.35", "10.99.70.37"}

	for _, host := range hosts {
		client := New(host, 22, user.Username, user.Password)
		wg.Add(1)
		go func(c *SSHCli) {
			for _, cmd := range []string{
				"pwd",
				"sss",
				"ps -ef|grep ansible",
				"whoami",
				"ifconfig",
				"/bin/bash -c 'ifconfig'",
				"/bin/bash --login -c 'ifconfig'",
				"sudo -H -S -n -u admin /bin/bash -c 'whoami && pwd && ifconfig'",
			} {
				run(c, cmd)
			}
			wg.Done()
		}(client)
	}
	wg.Wait()
}

func run(client *SSHCli, cmd string) {
	client.Run(cmd)
	fmt.Printf("Host:%s, cmd:%s, Result:%s, ExitCode:%d\n",
		fmt.Sprintf("%s:%d", client.Host, client.Port), client.Session.LastCmd, client.Session.LastResult, client.Session.ExitCode)

}
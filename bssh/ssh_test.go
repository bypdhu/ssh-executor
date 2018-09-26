package bssh

import (
	"sync"
	"testing"
	"fmt"
)

func TestConnect(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := New("10.99.70.38", 22, username, password)

	for _, cmd := range []string{"pwd", "sss", "ps -ef|grep ansible", "md5sum /tmp/bian/测试.txt"} {
		runCmdTest(client, cmd)
	}
}

func TestSSHCli_Run(t *testing.T) {
	var (
		wg sync.WaitGroup
	)
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	hosts := []string{"10.99.70.38", "10.99.70.34", "10.99.70.35", "10.99.70.37"}

	for _, host := range hosts {
		client := New(host, 22, username, password)
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
				runCmdTest(c, cmd)
			}
			wg.Done()
		}(client)
	}
	wg.Wait()
}

func runCmdTest(client *SSHCli, cmd string) {
	client.Run(cmd)
	fmt.Printf("Host:%s, cmd:%s, Result:%s, ExitCode:%d\n",
		fmt.Sprintf("%s:%d", client.Host, client.Port), client.Session.Cmd, client.Session.Result, client.Session.ExitCode)

}
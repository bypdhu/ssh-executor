package bssh

import (
	"testing"
	"fmt"

	"github.com/bypdhu/ssh-executor/utils"
)

func TestSFTPCli_PullFile(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	r := "/tmp/bian/mysql_slow.log"
	cp, _ := utils.GetFullPath(".")
	l := cp + "/../tmp/mysql_slow.txt"
	err := client.CopyOneFile("pull", r, l)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}

}
func TestSFTPCli_PullFile2(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	r := "/tmp/bian/aa.tar.gz"
	cp, _ := utils.GetFullPath(".")
	l := cp + "/../tmp/aa.tar.gz"
	err := client.CopyOneFile("pull", r, l)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}

}

func TestSFTPCli_PushFile(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	r := "/tmp/bian/config.yml"
	cp, _ := utils.GetFullPath(".")
	l := cp + "/../tmp/config.yml"

	err := client.CopyOneFile("push", l, r)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
}

func TestSFTPCli_PushFile2(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	r := "/tmp/bian/bb.tar.gz"
	cp, _ := utils.GetFullPath(".")
	l := cp + "/../tmp/bb.tar.gz"

	err := client.CopyOneFile("push", l, r)
	if err != nil {
		fmt.Printf("Error:%s", err)
	}
}
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

	cp, _ := utils.GetFullPath(".")

	r := "/tmp/bian/mysql_slow.txt"
	l := cp + "/../tmp/testdir"

	change, err := client.CopyOneFile(SftpPull, r, l)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	change, err = client.CopyOneFile(SftpPull, r, l)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	l2 := cp + "/../tmp/mysql_slow.38.txt"
	change, err = client.CopyOneFile(SftpPull, r, l2)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	r3 := "/tmp/bian/aa.tar.gz"
	l3 := cp + "/../tmp/aa.38.tar.gz"

	change, err = client.CopyOneFile(SftpPull, r3, l3)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	client.ForceCopy = true

	change, err = client.CopyOneFile(SftpPull, r3, l3)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	r4 := "/tmp/bian/aabb.tar.gz"
	l4 := cp + "/../tmp/aa.38.tar.gz"

	change, err = client.CopyOneFile(SftpPull, r4, l4)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

}

func TestSFTPCli_PushFile(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	cp, _ := utils.GetFullPath(".")

	r := "/tmp/bian/config.copy.yml"
	l := cp + "/../tmp/config.yml"

	change, err := client.CopyOneFile(SftpPush, l, r)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	r2 := "/tmp/bian/jumpserver.tar.gz"
	l2 := cp + "/../tmp/jumpserver.tar.gz"

	change, err = client.CopyOneFile(SftpPush, l2, r2)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

	client.ForceCopy = true
	change, err = client.CopyOneFile(SftpPush, l2, r2)
	fmt.Printf("Change:%s, Error:%s\n", change, err)

}

func TestSFTPCli_CopyManyFilesPush(t *testing.T) {
	cp, _ := utils.GetFullPath(".")
	srcDests := []*SrcDest{
		{Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml"},
		{Src:cp + "/../tmp/bb.tar.gz", Dest:"/tmp/bian/bb.copy.tar.gz"},
		{Src:cp + "/../tmp/mysql_slow.txt", Dest:"/tmp/bian/mysql_slow.copy.txt"},
		{Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
		{Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
	}

	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	client.IgnoreErr = true

	fmt.Printf("%s\n", srcDests)
	client.CopyManyFiles(SftpPush, srcDests)
	for _, sd := range srcDests {
		fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
	}
}

func TestSFTPCli_Run_Push(t *testing.T) {
	cp, _ := utils.GetFullPath(".")
	srcDests := []*SrcDest{
		{Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml"},
		{Src:cp + "/../tmp/bb.tar.gz", Dest:"/tmp/bian/bb.copy.tar.gz"},
		{Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
		{Src:cp + "/../tmp/mysql_slow.txt", Dest:"/tmp/bian/mysql_slow.copy.txt"},
		{Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
	}

	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	client.IgnoreErr = true
	client.ForceCopy = true

	fmt.Printf("%s\n", srcDests)
	client.Run(SftpPush, srcDests)
	for _, sd := range srcDests {
		fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
	}
}

func TestSFTPCli_Run_Pull(t *testing.T) {
	cp, _ := utils.GetFullPath(".")
	srcDests := []*SrcDest{
		{Dest:cp + "/../tmp/config.38.yml", Src:"/tmp/bian/config.yml"},
		{Dest:cp + "/../tmp/bb.38.tar.gz", Src:"/tmp/bian/bb.tar.gz"},
		{Dest:cp + "/../tmp/testdir1", Src:"/tmp/bian/"},
		{Dest:cp + "/../tmp/mysql_slow.38.txt", Src:"/tmp/bian/mysql_slow.txt"},
	}

	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	client.IgnoreErr = true
	client.ForceCopy = true

	fmt.Printf("%s\n", srcDests)
	client.Run(SftpPull, srcDests)
	for _, sd := range srcDests {
		fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
	}
}
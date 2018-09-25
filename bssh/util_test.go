package bssh

import (
	"testing"
	"fmt"
)

func TestGetRemoteFileMd5(t *testing.T) {
	username, password := getUser()
	fmt.Printf("username:%s, password:%s\n", username, password)

	client := NewSftp("10.99.70.38", 22, username, password)

	m, err := getRemoteFileMd5(client, "/tmp/bian/test.txt")
	fmt.Printf("md5:%s, err:%s, result:%s", m, err, client.Session.LastResult)

	m, err = getRemoteFileMd5(client, "/tmp/bian/测试.txt")
	fmt.Printf("md5:%s, err:%s, result:%s", m, err, client.Session.LastResult)
}

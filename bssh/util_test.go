package bssh

import (
    "testing"
    "fmt"
)

func TestGetRemoteFileMd5(t *testing.T) {
    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    c, _ := NewSftp("10.99.70.38", 22, username, password)

    m, err := c.GetRemoteFileMd5("/tmp/bian/test.txt")
    fmt.Printf("md5:%s, err:%s, result:%s", m, err, c.Stdout)

    m, err = c.GetRemoteFileMd5("/tmp/bian/测试.txt")
    fmt.Printf("md5:%s, err:%s, result:%s", m, err, c.Stdout)

    m, err = c.GetRemoteFileMd5("/tmp/bian/")
    fmt.Printf("md5:%s, err:%s, result:%s", m, err, c.Stdout)
}

func TestSFTPFileMode(t *testing.T) {

    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    c, _ := NewSftp("10.99.70.38", 22, username, password)

    fi, _ := c.sftpClient.Stat("/tmp/bian/test")
    fmt.Printf("%+v\n", fi)
    fmt.Printf("%+v\n", fi.Mode())
    fmt.Printf("%d\n", fi.Mode().Perm())
    fmt.Printf("%t\n", fi.Mode().IsDir())

}
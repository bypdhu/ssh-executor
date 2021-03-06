package bssh

import (
    "testing"
    "fmt"

    "github.com/bypdhu/ssh-executor/utils"
    "github.com/bypdhu/ssh-executor/common"
    "github.com/bypdhu/ssh-executor/task"
)

func TestSFTPCli_PullFile(t *testing.T) {
    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    client, _ := NewSftp("10.99.70.38", 22, username, password)

    cp, _ := utils.GetFullPath(".")

    for _, one := range []*task.CopyOneFile{
        {Src:"/tmp/bian/mysql_slow.txt", Dest:cp + "/../tmp/testdir"},
        {Src:"/tmp/bian/mysql_slow.txt", Dest:cp + "/../tmp/mysql_slow.38.txt"},
        {Src:"/tmp/bian/aa.tar.gz", Dest:cp + "/../tmp/aa.38.tar.gz"},
        {Src:"/tmp/bian/aa.tar.gz", Dest:cp + "/../tmp/aa.38.tar.gz", ForceCopy:true},
        {Src:"/tmp/bian/aabb.tar.gz", Dest:cp + "/../tmp/aa.38.tar.gz"},
        {Src:"/tmp/bian/", Dest:cp + "/../tmp/testdir1"},
    } {

        _ = client.CopyOneFile(common.SFTP_PULL.String(), one)
        fmt.Printf("Change:%s, Error:%s\n", one.Changed, one.Err)
    }

}

func TestSFTPCli_PushFile(t *testing.T) {
    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    client, _ := NewSftp("10.99.70.38", 22, username, password)
    client.Task = &task.Task{}

    cp, _ := utils.GetFullPath(".")

    for _, one := range []*task.CopyOneFile{
        {Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml"},
        {Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml", ForceCopy:true},
        {Src:cp + "/../tmp/jumpserver.tar.gz", Dest:"/tmp/bian/jumpserver.tar.gz"},
    } {
        _ = client.CopyOneFile(common.SFTP_PUSH.String(), one)
        fmt.Printf("Change:%s, Error:%s\n", one.Changed, one.Err)
    }

}

func TestSFTPCli_CopyManyFilesPush(t *testing.T) {
    cp, _ := utils.GetFullPath(".")
    srcDests := []*task.CopyOneFile{
        {Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml"},
        {Src:cp + "/../tmp/bb.tar.gz", Dest:"/tmp/bian/bb.copy.tar.gz"},
        {Src:cp + "/../tmp/mysql_slow.txt", Dest:"/tmp/bian/mysql_slow.copy.txt", ForceCopy:true},
        {Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
        {Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
    }

    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    client, _ := NewSftp("10.99.70.38", 22, username, password)
    client.Task = &task.Task{}

    client.IgnoreErr = true

    fmt.Printf("%+v\n", client)
    client.CopyArgs = task.CopyArgs{CopyFiles:srcDests, SftpMode:common.SFTP_PUSH.String()}
    fmt.Printf("%+v\n", client)

    client.SftpStart()
    for _, sd := range srcDests {
        fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
    }
}

func TestSFTPCli_Run_Push(t *testing.T) {
    cp, _ := utils.GetFullPath(".")
    srcDests := []*task.CopyOneFile{
        {Src:cp + "/../tmp/config.yml", Dest:"/tmp/bian/config.copy.yml"},
        {Src:cp + "/../tmp/bb.tar.gz", Dest:"/tmp/bian/bb.copy.tar.gz"},
        {Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
        {Src:cp + "/../tmp/mysql_slow.txt", Dest:"/tmp/bian/mysql_slow.copy.txt"},
        {Src:cp + "/../tmp/testdir", Dest:"/tmp/bian/testdir"},
    }

    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    client, _ := NewSftp("10.99.70.38", 22, username, password)

    client.Task = &task.Task{}
    client.IgnoreErr = true

    fmt.Printf("%+v\n", client)
    client.CopyArgs = task.CopyArgs{CopyFiles:srcDests, SftpMode:common.SFTP_PUSH.String()}
    fmt.Printf("%+v\n", client)

    client.SftpStart()
    for _, sd := range srcDests {
        fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
    }
}

func TestSFTPCli_Run_Pull(t *testing.T) {
    cp, _ := utils.GetFullPath(".")
    srcDests := []*task.CopyOneFile{
        {Dest:cp + "/../tmp/config.38.yml", Src:"/tmp/bian/config.yml"},
        {Dest:cp + "/../tmp/bb.38.tar.gz", Src:"/tmp/bian/bb.tar.gz"},
        {Dest:cp + "/../tmp/testdir1", Src:"/tmp/bian"},
        {Dest:cp + "/../tmp/mysql_slow.38.txt", Src:"/tmp/bian/mysql_slow.txt"},
    }

    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    client, _ := NewSftp("10.99.70.38", 22, username, password)
    client.Task = &task.Task{}

    client.IgnoreErr = true

    fmt.Printf("%+v\n", client)
    client.CopyArgs = task.CopyArgs{CopyFiles:srcDests, SftpMode:common.SFTP_PULL.String()}
    fmt.Printf("%+v\n", client)
    for _, cf := range client.CopyFiles {
        fmt.Printf("%+v\n", cf)
    }

    client.SftpStart()
    for _, sd := range srcDests {
        fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
    }
}

func TestSFTPCli_Run_Push_CreateDir(t *testing.T) {

    cp, _ := utils.GetFullPath(".")
    srcDests := []*task.CopyOneFile{{Src:cp + "/../tmp/config.yml",
        Dest:"/tmp/bian/test1/config.copy.yml",
        Owner:"admin", Group:"admin", ForceCopy:true,
        CreateDirectory:true, DirectoryMode:"777"}}

    username, password := getUser()
    //fmt.Printf("username:%s, password:%s\n", username, password)

    c, _ := NewSftp("10.99.70.38", 22, username, password)
    c.Task = &task.Task{}

    c.IgnoreErr = true

    fmt.Printf("%+v\n", c)
    c.CopyArgs.CopyFiles = srcDests
    c.CopyArgs.SftpMode = common.SFTP_PUSH.String()
    c.CopyArgs.Become = true
    fmt.Printf("%+v\n", c)
    for _, i := range c.CopyFiles {
        fmt.Printf("%+v\n", i)
    }

    c.SftpStart()
    for _, sd := range srcDests {
        fmt.Printf("changed:%t, %s\n", sd.Changed, sd.Err)
    }

}

func TestSFTPCli_ChownRemote(t *testing.T) {
    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    //c := NewSftp("10.99.70.38", 22, username, password)

    //err := c.ChownRemote("/tmp/bian/", "/tmp/bian/aa.tar.gz", "admin", "admin")
    //fmt.Printf("err:%s, result:%s", err, c.Result)
}

func TestSFTPCli_GetDirExists(t *testing.T) {
    username, password := getUser()
    fmt.Printf("username:%s, password:%s\n", username, password)

    c, _ := NewSftp("10.99.70.38", 22, username, password)

    base, create := c.GetDirExists("/tmp/bian/test/testcc")
    fmt.Printf("base:%s,create:%s\n\n\n", base, create)

    base, create = c.GetDirExists("/tmp/bian/test/test2/testcc/")
    fmt.Printf("base:%s,create:%s\n\n\n", base, create)

    base, create = c.GetDirExists("/tmp/bian/test/test2/testcc")
    fmt.Printf("base:%s,create:%s\n\n\n", base, create)

    base, create = c.GetDirExists("/tmp/bian/test/")
    fmt.Printf("base:%s,create:%s\n\n\n", base, create)

    base, create = c.GetDirExists("/tmp/bian/test")
    fmt.Printf("base:%s,create:%s\n\n\n", base, create)

}
package bssh

import (
    "fmt"
    "net"
    "time"
    "golang.org/x/crypto/ssh"
)

type Cli struct {
    Host     string
    Port     int
    Username string
    Password string
    client   *ssh.Client
    SSHConfig
}

type SSHConfig struct {
    Timeout int
}

func (c *Cli) newClient() (err error) {
    var (
        auth []ssh.AuthMethod
        config ssh.Config
        clientConfig *ssh.ClientConfig
    )
    auth = make([]ssh.AuthMethod, 0)
    auth = append(auth, ssh.Password(c.Password))

    config = ssh.Config{
        Ciphers:[]string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128",
            "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
    }
    clientConfig = &ssh.ClientConfig{
        User: c.Username,
        Auth: auth,
        Timeout: time.Duration(c.SSHConfig.Timeout),
        Config: config,
        HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
    }

    c.client, err = ssh.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port), clientConfig)
    if err != nil {
        return
    }

    return
}

func (c *Cli) closeClient() (err error) {
    if c.client != nil {
        err = c.client.Close()
        c.client = nil
    }
    return
}


# ssh-executor
ssh executor implement by go.


## Launch

### help
    usage: ssh-executor[.exe] [<flags>]
    
    ssh-executor
    
    Flags:
      -h, --help                   Show context-sensitive help (also try --help-long
                                   and --help-man).
          --version                Show application version.
          --log.level="info"       Only log messages with the given severity or
                                   above. Valid levels: [debug, info, warn, error,
                                   fatal]
          --log.format="logger:stderr"
                                   Set the log target and format. Example:
                                   "logger:syslog?appname=bob&local=7" or
                                   "logger:stdout?json=true"
      -c, --config.file=""         application's configuration file path.
      -T, --launch.type="direct"   server/direct;default direct. server will setup a
                                   http server. direct will execute command once.
      -t, --ssh.timeout=30         timeout in ssh connection. default 30s.
          --web.listen_address=""  [launch.type=server] Address to listen on for UI,
                                   API.
          --telemetry.listen_address=""
                                   [launch.type=server] Address to listen on for
                                   telemetry.
      -i, --hosts=""               [launch.type=direct] Hosts to connect by ssh.
                                   Combined by ','. Add hosts.file.
      -f, --hosts.file=""          [launch.type=direct] File of hosts to connect by
                                   ssh. One ip on line. Add hosts.
      -u, --user.name=""           [launch.type=direct] Username for ssh connection.
      -p, --user.pass=""           [launch.type=direct] Password for ssh connection.
      -C, --command=""             [launch.type=direct] Command for ssh connection.

### example

#### use as a tool.
    
    linux 
        chmod +x ssh-executor
        ./ssh-excutor command
    windows
        ssh-executor.exe command
    
    ./ssh-executor -i 10.99.70.35,10.99.70.38 -u user --user.pass=pass -C "/bin/sh --login -c 'ifconfig'"


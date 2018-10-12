# ssh-executor

| Date      | Version | Owner | Comments    |
| --------- | ------- | ----- | ----------- |
| 2018.10.12 | v1.0.0 | bian  | Created it. |
|           |         |       |             |

## 一、简介
使用go语言实现的ssh命令执行器。

- 提供了两种调用方式。
    - 一次性启动并执行。
    - 启动为http服务。
- 实现了两个模块，包括shell和copy。
    - shell: 直接执行shell命令
    - copy: 远程copy文件

## 二、启动

### 1. 下载

### 2. 帮助信息
    usage: ssh-executor[.exe] [<flags>]
    
    ssh-executor
    
    Flags:
      -h, --help                   Show context-sensitive help (also try --help-long and --help-man).
          --version                Show application version.
          --log.level="info"       Only log messages with the given severity or above. Valid levels: [debug, info, warn, error, fatal]
          --log.format="logger:stderr"  
                                   Set the log target and format. Example: "logger:syslog?appname=bob&local=7" or "logger:stdout?json=true"
      -c, --config.file=""         application's configuration file path.
      -T, --launch.type="DIRECT"   server/direct;default direct. server will setup a http server. direct will execute command once.
      -t, --ssh.timeout=30         timeout in ssh connection. default 30s.
      -a, --web.listen_address=""  [launch.type=server] Address to listen on for UI, API.
          --telemetry.listen_address=""  
                                   [launch.type=server] Address to listen on for telemetry.
      -i, --hosts=""               [launch.type=direct] Hosts to connect by ssh. Combined by ','. Add hosts.file.
      -f, --hosts.file=""          [launch.type=direct] File of hosts to connect by ssh. One ip on line. Add hosts.
      -u, --user.name=""           [launch.type=direct] Username for ssh connection.
      -p, --user.pass=""           [launch.type=direct] Password for ssh connection.
      -m, --module="SHELL"         [launch.type=direct] Module to handle. like 'shell' 'copy'
      -C, --command=""             [launch.type=direct] Command to handle.

### 3. 启动举例

#### 3.1 一次性启动并执行
    
##### 3.1.1 简单启动并执行

- linux
    
     ```
     chmod +x ssh-executor 
     ./ssh-executor -i 10.99.70.35,10.99.70.38 -u user -p pass -C "/bin/sh --login -c 'ifconfig'"
    
     ```
- windows

    ```
    ssh-executor.exe command
    ```
    
##### 3.1.2 其他参数启动并执行

    ./ssh-executor -i 10.99.70.35,10.99.70.38 -f host.txt -u user -p pass -C "/bin/sh --login -c 'ifconfig'" -t 10

#### 3.2 作为http服务启动

    ./ssh-executor -T server -a localhost:9888
    
## 三、使用

### 1. 一次性启动使用

#### 1.1 简单命令调用
    ./ssh-executor -i 10.99.70.35,10.99.70.38 -u user -p pass -C "/bin/sh --login -c 'ifconfig'"

#### 1.2 使用yaml文件执行task

### 2. 作为服务启动调用

#### 2.1 执行shell和copy任务

**请求语法**
```
POST /job HTTP/1.1 
Content-Type: application/json
```
```json
{
  "user_flag": "test",
  "hosts": [
    "10.99.70.38",
    "10.99.70.35"
  ],
  "ssh_config": {
    "timeout": 30,
    "sh": "/bin/sh",
    "username": "user",
    "password": "pass"
  },
  "tasks": [
    {
      "name": "ddd",
      "module": "shell",
      "args": {
        "shell": {
          "command": "ls",
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
          "sftp_mode": "pull",
          "copy_files": [
            {
              "src": "/tmp/foo.yml",
              "dest": "foo.yml",
              "owner": "foo",
              "group": "foo",
              "mode": 644,
              "md5": "12313131",
              "force": true,
              "create_directory": true,
              "recursive": true,
              "directory_mode": 755
            }
          ]
        }
      }
    }
  ]
}
```

**返回体**
```
Content-Type: application/json
示例，与请求不对应，包括任务的成功和失败情况

```
```json
{
    "success": false,
    "msg": "",
    "detail": [
        {
            "host": "10.99.70.38",
            "success": false,
            "tasks": [
                {
                    "name": "shell error task",
                    "module": "shell",
                    "success": false,
                    "result": {
                        "stdout": "ls: cannot open directory /home/abcedfg: Permission denied\r\n",
                        "exitcode": 1
                    },
                    "err": "Process exited with status 1"
                },
                {
                    "name": "shell success task",
                    "module": "shell",
                    "success": true,
                    "result": {
                        "stdout": "admin\r\n",
                        "exitcode": 0
                    },
                    "err": ""
                },
                {
                    "name": "copy task",
                    "module": "copy",
                    "sftpmode": "push",
                    "success": false,
                    "copyfiles": [
                        {
                            "src": "tmp/config.yml",
                            "dest": "/tmp/bian/config.web.yml",
                            "result": {
                                "changed": true
                            },
                            "err": ""
                        },
						{
                            "src": "tmp/config.notexist.yml",
                            "dest": "/tmp/bian/config.web.yml",
                            "result": {
                                "changed": false
                            },
                            "err": "open tmp/config.notexist.yml: The system cannot find the file specified."
                        }
                    ]
                }
            ]
        }
    ]
}
```


## 更新历史

### v1.0.0
- 首次提交
- 支持直接调用执行shell命令
- 支持直接调用，读取文件的任务执行功能模块
- 支持作为http服务启动，调用执行功能模块
- 功能模块
    - shell: 直接执行shell命令
    - copy: 远程copy文件
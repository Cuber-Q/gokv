#gokv

gokv是一个分布式K-V存储系统，类似[etcd](https://github.com/coreos/etcd)

以GO语言编写，通过命令行进行操作

##1、命令简介

####1.1 服务端
gokv 根命令帮助

`$ gokv`
```bash
gokv is a simple fast k-v store system coded in go

Usage:
  gokv [command]

Available Commands:
  help        Help about any command
  server      start gokv server

Flags:
  -h, --help   help for gokv

Use "gokv [command] --help" for more information about a command.
```

gokv server 子命令帮助

`$ gokv server`
```bash
start gokv server

Usage:
  gokv server [flags]

Flags:
  -h, --help       help for server
  -P, --port int   specify server port (default 9901)
```

####1.2 客户端

gokv-cli 客户端命令行简介

`$ gokv-cli -h`
```bash
gokv-cli is a simple fast k-v store system command app coded in go

Usage:
  gokv-cli [flags]
  gokv-cli [command]

Available Commands:
  get         get value of <key>
  help        Help about any command
  set         set <key> <value>

Flags:
      --endpoint string   specify endponits (default "127.0.0.1:9901")
  -h, --help              help for gokv-cli

Use "gokv-cli [command] --help" for more information about a command.

```
启动客户端，连接到服务端：

`$ gokv-cli`
```bash
2018/07/11 22:16:52 current endpoints are:
 127.0.0.1:9901
gokv-cli>
```
客户端命令行使用[go-prompt](https://github.com/c-bata/go-prompt)，提供交互式命令

#####1.2.1 存储

```bash
gokv-cli>set mykey myvalue
OK
```
#####1.2.2 查询
```bash
gokv-cli>get mykey
myvalue
```
其他命令待补充...

##2、集群模式(待补充)

###2.1 主从选举

###2.2 数据同步

...
# TiM

TiM is a tool for managing multiple tidb clusters. This tool currently needs to be work with tidb-ansible.  
Using this tool can greatly improve the efficiency of tidb ops and greatly reduce the issue of cluster failure caused by incorrect operation.

## Installation

build a binary from the source code, run the following commands. 
Note that compiling requires Go version 1.12+.

```$go
$ git clone https://github.com/tidbops/tim.git
$ cd tim
$ make
```

## Usage 

```shell
./bin/tim help
A tool to manage multi tidb-ansible and help to upgrade tidb version

Usage:
  tim [command]

Available Commands:
  env         init environment for tidb-ansible
  help        Help about any command
  init        init tidb-ansible files
  list        tidb-clusters list info
  search      tidb-clusters search info
  upgrade     upgrade tidb version, just generate the new version tidb-ansible files

Flags:
  -d, --detach          Run ctl without readline. (default true)
  -h, --help            Help message.
  -i, --interact        Run tim with readline.
  -L, --level string    log level, support info / warning / debug / error / fatal (default "info")
  -u, --server string   tim-server address
  -V, --version         Print version information and exit.

Use "tim [command] --help" for more information about a command.
```

### How to rolling update?

*ps: currently only supports automatic generation of **tikv config file***

```$xslt
upgrade tidb version, just generate the new version tidb-ansible files

Usage:
  tim upgrade <name> [flags]

Flags:
      --rule-file string        rule files for different version of configuration conversion
      --target-version string   the version that ready to upgrade to

Global Flags:
  -d, --detach          Run ctl without readline. (default true)
  -h, --help            Help message.
  -i, --interact        Run tim with readline.
  -L, --level string    log level, support info / warning / debug / error / fatal (default "info")
  -u, --server string   The tim-server address
  -V, --version         Print version information and exit.
```

* prepare rule file 

`@new` for adding new config in target version,  
`@delete` for deleting config from origin config
eg:  

```yaml
# @new
---
pessimistic_txn:

rocksdb:
  titan:
  defaultcf:
    titan:

storage:
  block-cache:

# @delete
---
delete:
  - "storage"
``` 

### Demo

![Demo](https://github.com/tidbops/tim/blob/master/images/demo.gif?raw=true)

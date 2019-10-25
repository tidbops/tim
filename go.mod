module github.com/tidbops/tim

go 1.12

require (
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/imdario/mergo v0.3.8
	github.com/kylelemons/godebug v1.1.0
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/manifoldco/promptui v0.3.2
	github.com/mattn/go-isatty v0.0.10 // indirect
	github.com/mattn/go-shellwords v1.0.6
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/ngaut/log v0.0.0-20180314031856-b8e36e7ba5ac
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c // indirect
	gopkg.in/mikefarah/yaml.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.2
	xorm.io/core v0.7.2
	xorm.io/xorm v0.8.0
)

replace gopkg.in/mikefarah/yaml.v2 => github.com/mikefarah/yaml/v2 v2.4.0

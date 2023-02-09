module github.com/tidbops/tim

go 1.12

require (
	github.com/bndr/gotabulate v1.1.2
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e
	github.com/gin-gonic/gin v1.7.7
	github.com/jessevdk/go-assets v0.0.0-20160921144138-4f4301a06e15
	github.com/kylelemons/godebug v1.1.0
	github.com/logrusorgru/aurora v0.0.0-20191017060258-dc85c304c434
	github.com/manifoldco/promptui v0.3.2
	github.com/mattn/go-shellwords v1.0.6
	github.com/mattn/go-sqlite3 v1.10.0
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ngaut/log v0.0.0-20180314031856-b8e36e7ba5ac
	github.com/nicksnyder/go-i18n v1.10.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/tidbops/mergo v0.3.9-0.20191026204819-85b7c6436930
	google.golang.org/appengine v1.6.5 // indirect
	gopkg.in/alecthomas/kingpin.v3-unstable v3.0.0-20180810215634-df19058c872c // indirect
	gopkg.in/mikefarah/yaml.v2 v2.4.0
	gopkg.in/yaml.v2 v2.2.8
	xorm.io/core v0.7.2
	xorm.io/xorm v0.8.0
)

replace gopkg.in/mikefarah/yaml.v2 => github.com/mikefarah/yaml/v2 v2.4.0

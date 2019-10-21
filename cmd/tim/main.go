package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/chzyer/readline"
	shellwords "github.com/mattn/go-shellwords"
	flag "github.com/spf13/pflag"
	"github.com/tidbops/tim/pkg/ctl"
	v "github.com/tidbops/tim/pkg/version"
)

var (
	url      string
	detach   bool
	interact bool
	version  bool
	help     bool
)

func init() {
	flag.StringVarP(&url, "server", "u", "", "The tim-server address")
	flag.BoolVarP(&detach, "detach", "d", true, "Run ctl without readline.")
	flag.BoolVarP(&interact, "interact", "i", false, "Run tim with readline.")
	flag.BoolVarP(&version, "version", "V", false, "Print version information and exit.")
	flag.BoolVarP(&help, "help", "h", false, "Help message.")
}

func main() {
	timAddr := os.Getenv("TIM_SERVER")
	if timAddr != "" {
		os.Args = append(os.Args, "-u", timAddr)
	}

	flag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		v.PrintVersionInfo()
		os.Exit(0)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		fmt.Printf("\nGot signal [%v] to exit.\n", sig)
		switch sig {
		case syscall.SIGTERM:
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}()
	var input []string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		detach = true
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			return
		}
		input = strings.Split(strings.TrimSpace(string(b[:])), " ")
	}
	if interact {
		loop()
		return
	}
	ctl.Start(append(os.Args[1:], input...))
}

func loop() {
	l, err := readline.NewEx(&readline.Config{
		Prompt:            "\033[31m»\033[0m ",
		HistoryFile:       "/tmp/readline.tmp",
		InterruptPrompt:   "^C",
		EOFPrompt:         "^D",
		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		line, err := l.Readline()
		if err != nil {
			if err == readline.ErrInterrupt {
				break
			} else if err == io.EOF {
				break
			}
			continue
		}
		if line == "exit" {
			os.Exit(0)
		}
		args, err := shellwords.Parse(line)
		if err != nil {
			fmt.Printf("parse command err: %v\n", err)
			continue
		}
		args = append(args, "-u", url)

		ctl.Start(args)
	}
}

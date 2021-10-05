package main

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/vic020/go-crawler/core/manager"
)

func daemonize(args ...string) {

	var newArgs []string

	if len(args) > 1 {
		newArgs = args[1:]
	}
	cmd := exec.Command(args[0], newArgs...)
	cmd.Env = os.Environ()
	cmd.Start()

}

func isDaemonized() {
	daemon := false

	args := os.Args

	for k, v := range args {
		if v == "-d" {
			daemon = true
			args[k] = ""
		}
	}

	if daemon {
		daemonize(args...)
		syscall.Exit(0)
	}
}

func main() {
	isDaemonized()

	manager.GetInstance().Run()
}

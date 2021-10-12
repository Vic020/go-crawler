package utils

import (
	"os"
	"os/exec"
	"syscall"

	"github.com/vic020/go-crawler/conf"
	"github.com/vic020/go-crawler/utils/logger"
)

func daemonize() {

	args := os.Args

	// remove -d flag
	for k, v := range args {
		if v == "-d" {
			args[k] = ""
		}
	}

	var newArgs []string

	if len(args) > 1 {
		newArgs = args[1:]
	}
	cmd := exec.Command(args[0], newArgs...)
	cmd.Env = os.Environ()
	cmd.Start()

}

func IsDaemonized() {
	logger.Info("Is Daemonized: ", conf.DaemonMode)

	if !conf.DaemonMode {
		return
	}

	daemonize()
	syscall.Exit(0)
}

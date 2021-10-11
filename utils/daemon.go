package utils

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
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

func IsDaemonized() {
	daemon := false

	args := os.Args

	for k, v := range args {
		if v == "-d" {
			daemon = true
			args[k] = ""
		}
	}

	fmt.Println("Is Daemonized: ", daemon)

	if daemon {
		daemonize(args...)
		syscall.Exit(0)
	}
}

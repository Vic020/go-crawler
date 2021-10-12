package conf

import (
	"flag"
	"fmt"
)

var (
	// log config
	LogPath    string
	DebugMode  bool
	DaemonMode bool

	//
	Commands []string
	Command  string
)

func usageFunc() {
	w := flag.CommandLine.Output()

	fmt.Fprintf(w, "Go Crawler is a web crawler system build by go.\n\n")

	fmt.Fprintf(w, "Usage:\n\n\tgo-crawler <command> [arguments]\n\n")

	fmt.Fprintf(w, "The commands:\n\tall\tlaunch all components\n\n")

	fmt.Fprintln(w, "The args:")
	flag.PrintDefaults()
}

func InitFlag() {
	flag.StringVar(&LogPath, "log", "logs/", "logs file")
	flag.BoolVar(&DebugMode, "debug", false, "debug mode")
	flag.BoolVar(&DaemonMode, "d", false, "daemon mode")
	flag.Parse()

	Commands = flag.Args()
	Command = flag.Arg(0)

	flag.Usage = usageFunc
}

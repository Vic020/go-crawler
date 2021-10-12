package main

import (
	"flag"

	"github.com/vic020/go-crawler/conf"
	"github.com/vic020/go-crawler/core/manager"
	"github.com/vic020/go-crawler/utils"
)

func init() {
	conf.InitFlag()
}

func main() {
	switch conf.Command {
	case "":
		flag.Usage()
		return
	case "all":
		utils.IsDaemonized()

		manager.GetInstance().Run()
	}
}

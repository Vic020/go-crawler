package main

import (
	"github.com/vic020/go-crawler/core/manager"
	"github.com/vic020/go-crawler/utils"
)

func main() {
	utils.IsDaemonized()

	manager.GetInstance().Run()
}

package main

import "github.com/vic020/go-crawler/api"

func main() {
	server := api.NewHTTPServer()

	server.Run(":8000")
}

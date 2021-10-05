package manager

import (
	"flag"
	"io"
	"log"
	"os"
	"path"
	"sync"
)

type Manager struct {
	id      string
	logPath string
}

var (
	manager *Manager
	once    sync.Once
)

func (c *Manager) initFlag() {
	flag.StringVar(&c.logPath, "log", "logs/", "logs file")
}

func (c *Manager) initLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	f, err := os.OpenFile(path.Join(c.logPath, "go-manager.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	log.SetOutput(io.MultiWriter(f, os.Stdout))
}

func GetInstance() *Manager {
	once.Do(func() {
		manager = &Manager{
			id: "vic",
		}
		manager.initFlag()
		manager.initLog()
	})

	return manager
}

func (c *Manager) Run() {
}

func (c *Manager) Close() {

}

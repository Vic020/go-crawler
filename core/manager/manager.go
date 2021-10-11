package manager

import (
	"flag"
	"fmt"
	"sync"

	"github.com/vic020/go-crawler/api"
	"github.com/vic020/go-crawler/core/fetcher"
	"github.com/vic020/go-crawler/core/limiter"
	"github.com/vic020/go-crawler/models"
	"github.com/vic020/go-crawler/utils/logger"
)

type Manager struct {
	id       string
	logPath  string
	services Services
}

type Services struct {
	fetcher    []*fetcher.Fetcher
	httpServer *api.HTTPServer
}

var (
	manager *Manager
	once    sync.Once
)

func (c *Manager) initFlag() {
	flag.StringVar(&c.logPath, "log", "logs/", "logs file")
}

func (c *Manager) initFetcher(num int, limiter *limiter.Limiter, in, out chan models.Task) {
	logger.Info("Fetcher initializing...")
	for i := 0; i < num; i++ {
		c.services.fetcher = append(c.services.fetcher, fetcher.NewFetcher(
			limiter,
			in,
			out,
		))
	}
	logger.Info("Fetcher initialized")
}

func (c *Manager) Init() {
	c.initFlag()
	logger.InitLogger(c.logPath)

	logger.Info("Log init completed")

	FetchTasks := make(chan models.Task, 100)
	ResultTasks := make(chan models.Task, 100)

	limiter := limiter.NewLimiter(10)

	c.initFetcher(1, limiter, FetchTasks, ResultTasks)

	c.services.httpServer = api.NewHTTPServer()

	logger.Info("Manager initialized")
}

func GetInstance() *Manager {
	once.Do(func() {
		fmt.Println("Manager initializing...")
		manager = &Manager{
			id: "vic",
		}
		manager.Init()
	})

	return manager
}

func (c *Manager) Run() {
	for _, fetch := range c.services.fetcher {
		fetch.Run()
	}

	c.services.httpServer.Run("0.0.0.0:8000")

}

func (c *Manager) Close() {
	c.services.httpServer.Close()
	for _, fetch := range c.services.fetcher {
		fetch.Close()
	}
}

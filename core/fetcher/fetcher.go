package fetcher

import (
	"time"

	"github.com/google/uuid"
	"github.com/levigross/grequests"
	"github.com/vic020/go-crawler/core/limiter"
	"github.com/vic020/go-crawler/models"
	"github.com/vic020/go-crawler/utils/logger"
)

type Fetcher struct {
	id        string           // fetcher uuid
	limiter   *limiter.Limiter // fetcher rate limit
	taskQueue chan models.Task // fetcher tasks queue
	outQueue  chan models.Task // fetcher result tasks queue
	signal    chan int         // fetcher signal
	close     chan int         // fetcher close signal
	isRunning bool             // fetcher status
}

func (fc *Fetcher) process(task models.Task) {
	logger.Info(fc.id, task.ID, task.URL, "fetch")
	// blocking for token
	fc.limiter.Take()

	logger.Info(fc.id, task.ID, task.URL, "fetching")

	rawHtml := fc.fetch(task)

	task.RawHTML = rawHtml

	logger.Info(fc.id, task.ID, task.URL, "fetched")

	fc.outQueue <- task

}

func (fc *Fetcher) fetch(task models.Task) string {

	r, err := grequests.Get(task.URL, nil)

	if err != nil {
		logger.Error(err)
	}

	return r.String()

}

func (fc *Fetcher) loop() {
	for {
		select {

		case singal := <-fc.signal:
			// control signal
			switch singal {
			case 0:
				logger.Infof("Fetcher %v, loop get close signal", fc.id)
				// Fetcher exit
				close(fc.signal)
				fc.close <- 0
				return
			}
		case task := <-fc.taskQueue:
			// blocking task process
			fc.process(task)
		default:
			// for mutiple fetcher run in empty, finally make goroutine fake dead
			time.Sleep(1 * time.Second)
		}
	}
}

func (fc *Fetcher) GetId() string {
	return fc.id
}

func NewFetcher(limiter *limiter.Limiter, inQueue, outQueue chan models.Task) *Fetcher {
	f := &Fetcher{
		id:        uuid.NewString(),
		limiter:   limiter,
		taskQueue: inQueue,
		outQueue:  outQueue,
		signal:    make(chan int, 1),
		close:     make(chan int),
		isRunning: false,
	}

	logger.Info("New fetch inited, id: ", f.id)

	return f

}

func (fc *Fetcher) Run() {
	if fc.isRunning {
		return
	}

	go fc.loop()
	fc.isRunning = true

	logger.Infof("Fetcher %v is running", fc.id)
}

func (fc *Fetcher) Close() {
	if !fc.isRunning {
		return
	}

	logger.Infof("Fetcher %v is closing", fc.id)

	fc.signal <- 0
	<-fc.close
	close(fc.close)
}

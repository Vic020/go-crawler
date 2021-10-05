package fetcher

import (
	"time"

	"github.com/google/uuid"
	"github.com/levigross/grequests"
	"github.com/vic020/go-crawler/models"
	"github.com/vic020/go-crawler/utils/logger"
)

type Fetcher struct {
	id        string           // fetcher uuid
	limiter   chan int         // fetcher rate limit
	taskQueue chan models.Task // fetcher tasks queue
	outQueue  chan models.Task // fetcher result tasks queue
	signal    chan int         // fetcher signal
	close     chan int         // fetcher close signal
	isRunning bool             // fetcher status
}

func (fc *Fetcher) process(task models.Task) {
	// blocking for token
	<-fc.limiter

	rawHtml := fc.fetch(task)

	task.RawHTML = rawHtml

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

func NewFetcher(limiter chan int, inQueue, outQueue chan models.Task) *Fetcher {
	return &Fetcher{
		id:        uuid.NewString(),
		limiter:   limiter,
		taskQueue: inQueue,
		outQueue:  outQueue,
		signal:    make(chan int, 1),
		close:     make(chan int),
		isRunning: false,
	}
}

func (fc *Fetcher) Run() {
	if fc.isRunning {
		return
	}
	go fc.loop()
	fc.isRunning = true
}

func (fc *Fetcher) Close() {
	if !fc.isRunning {
		return
	}

	fc.signal <- 0
	<-fc.close
	close(fc.close)
}

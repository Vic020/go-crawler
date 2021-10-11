package limiter

import (
	"os"
	"time"

	"github.com/vic020/go-crawler/utils/logger"
)

type Limiter struct {
	// config
	qps int

	// inner resources
	bucket    chan int
	signal    chan int
	close     chan int
	isRunning bool
}

func NewLimiter(qps int) *Limiter {
	if qps == 0 {
		logger.Error("qps = 0")
		os.Exit(1)
	}

	return &Limiter{
		qps:       qps,
		bucket:    make(chan int, qps),
		signal:    make(chan int, 1),
		close:     make(chan int),
		isRunning: false,
	}
}

func (l *Limiter) genToken(d time.Duration) {
	// time for generate token
	t := time.NewTicker(d)

	for {
		select {
		case <-t.C:
			// gen token main worker
			select {
			case l.bucket <- 1:
			default:
				continue
			}
		case signal := <-l.signal:
			// close control
			switch signal {
			case 0:
				close(l.signal)
				l.close <- 0
				return
			}
		}
	}

}

func (l *Limiter) Take() {
	<-l.bucket
}

func (l *Limiter) Run() {
	if l.isRunning {
		logger.Error("Limit is running")
		return
	}

	go l.genToken(time.Duration(1000/l.qps) * time.Millisecond)
	l.isRunning = true
}

func (l *Limiter) Close() {
	if !l.isRunning {
		return
	}

	l.signal <- 0
	<-l.close
	close(l.close)
	l.isRunning = false
}

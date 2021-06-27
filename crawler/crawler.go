package crawler

type Crawler struct {
	Tasks chan int
}

type CrawlerOptions struct {
	Id           int
	ExecInterval int
}

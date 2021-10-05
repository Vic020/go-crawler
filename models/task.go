package models

import (
	"encoding/json"
	"time"
)

type Task struct {
	// Base
	ID           string `json:"id"`
	CreatedTime  int64  `json:"created_time"`
	ModifiedTime int64  `json:"modified_time"`
	DeletedTime  int64  `json:"deleted_time"`

	// Crawl
	URL                string   `json:"url"`
	Path               string   `json:"path"`
	Query              string   `json:"query"`
	RawHTML            string   `json:"raw_html"`
	LastestCrawledTime int64    `json:"lastest_crawled_time"`
	CrawledTimes       []string `json:"crawled_times"`

	// Sechedule
	Crontab     string `json:"crontab"`
	CoolingTime string `json:"cooling_time"`

	// Result
	Selectors []SelectorAtom `json:"selectors"`
}

type SelectorAtom struct {
	CssSelect string `json:"css_select"`
	Xpath     string `json:"xpath"`
	Result    string `json:"result"`
}

func (t *Task) ToJson() string {
	b, err := json.Marshal(t)
	if err != nil {
		// logs or panic
		return "{}"
	}
	return string(b)
}

func (t *Task) modifyTime() {
	t.ModifiedTime = time.Now().Unix()
}

func (t *Task) SetRawHTML(raw string) {
	t.RawHTML = raw
	t.LastestCrawledTime = time.Now().Unix()
	t.modifyTime()
}

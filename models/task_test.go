package models_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/vic020/go-crawler/models"
)

func newTask() models.Task {

	task := models.Task{
		ID:                 uuid.NewString(),
		CreatedTime:        0,
		ModifiedTime:       0,
		DeletedTime:        0,
		URL:                "",
		Path:               "",
		Query:              "",
		RawHTML:            "",
		LastestCrawledTime: 0,
		CrawledTimes:       []int64{},
		Crontab:            "",
		CoolingTime:        "",
		Selectors: []models.SelectorAtom{
			{
				CssSelect: "",
				Xpath:     "",
				Result:    "",
			},
			{
				CssSelect: "",
				Xpath:     "",
				Result:    "",
			},
		},
	}
	return task
}

func TestTask(t *testing.T) {
	task := newTask()
	t.Log("task is", task)

	if "{}" == task.ToJson() {
		t.Error("Json Marshal gets wrong")
	}
}

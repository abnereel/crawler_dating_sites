package view

import (
	"os"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/frontend/model"
	common "github.com/abnereel/crawler_dating_sites/crawler/model"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	view := CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")

	page := model.SearchResult{}
	page.Hits = 123
	item := engine.Item{
		Url:  "http://www.hongniang.com/user/6/a/11011577.html",
		Type: "hongniang",
		Id:   "97046681",
		Payload: common.User{
			Name:      "安静的雪",
			Gender:    "女",
			Age:       "34岁",
			Height:    "162cm",
			Marriage:  "已婚",
			Education: "本科",
			WorkAdd:   "山东菏泽",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items, item)
	}

	err = view.Render(out, page)
	if err != nil {
		panic(err)
	}
}
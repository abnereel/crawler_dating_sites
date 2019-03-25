package parser

import (
	"fmt"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
	"testing"
)

func TestParseProfile(t *testing.T) {
	expected := model.Profile{
		Name:      "安静的雪",
		Gender:    "女",
		Age:       "34岁",
		Height:    "162cm",
		Marriage:  "已婚",
		Education: "本科",
		Address:   "山东菏泽",
	}

	fmt.Println(expected)
}

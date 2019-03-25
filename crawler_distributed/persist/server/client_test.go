package main

import (
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/model"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/config"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/rpcsupport"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)

	// start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	// Call save
	item := engine.Item{
		Url:  "http://www.hongniang.com/user/6/a/11011577.html",
		Type: "hongniang",
		Id:   "97046681",
		Payload: model.Profile{
			Name:      "安静的雪",
			Gender:    "女",
			Age:       "34岁",
			Height:    "162cm",
			Marriage:  "已婚",
			Education: "本科",
			Address:   "山东菏泽",
		},
	}

	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}

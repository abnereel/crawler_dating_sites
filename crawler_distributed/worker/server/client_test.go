package main

import (
	"fmt"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/config"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/rpcsupport"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/worker"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServerRpc(host, worker.CrawlService{})

	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	req := worker.Request{
		Url: "http://www.hongniang.com/user/member/id/10705401",
		Parser: worker.SerializedParser{
			Name: config.ParseUser,
			Args: "",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(result)
	}
}
package main

import (
	"flag"
	"log"
	"net/rpc"
	"strings"
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/scheduler"
	"github.com/abnereel/crawler_dating_sites/crawler/zghongniang/parser"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/config"
	itemSaver "github.com/abnereel/crawler_dating_sites/crawler_distributed/persist/client"
	"github.com/abnereel/crawler_dating_sites/crawler_distributed/rpcsupport"
	worker "github.com/abnereel/crawler_dating_sites/crawler_distributed/worker/client"
)

var (
	itemsaverHost = flag.String("itemsaverhost", "", "itemsaverhost", )

	workerHosts = flag.String("workerhosts", "", "workerhosts (comma separated)")
)

func main() {
	flag.Parse()
	itemChan, err := itemSaver.ItemSaver(*itemsaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))

	processor := worker.CreateProcessor(pool)

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	e.Run2(engine.Request{
		Url:    "http://www.hongniang.com/index/search?&page=",
		Parser: engine.NewFuncParser(parser.ParseUserList, config.ParseUserList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpcsupport.NewClient(h)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", h)
		} else {
			log.Printf("error connecting to %s: %v", h, err)
		}
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()

	return out
}

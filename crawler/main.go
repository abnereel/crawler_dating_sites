package main

import (
	"github.com/abnereel/crawler_dating_sites/crawler/engine"
	"github.com/abnereel/crawler_dating_sites/crawler/persist"
	"github.com/abnereel/crawler_dating_sites/crawler/scheduler"
	"github.com/abnereel/crawler_dating_sites/crawler/zghongniang/parser"
	_ "github.com/abnereel/crawler_dating_sites/crawler/zhenai/parser"
)

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	/*engine.Run2(engine.Request{
		Url:        "http://www.hongniang.com/index/search?&page=",
		ParserFunc: parser2.ParseUserList,
	})*/

	itemChan, err := persist.ItemSaver("dating_user")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
		RequestProcessor: engine.Worker,
	}
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/
	/*e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/
	e.Run2(engine.Request{
		Url:        "http://www.hongniang.com/index/search?&page=",
		Parser: engine.NewFuncParser(parser.ParseUserList, "ParseUserList"),
	})
}

package main

import (
	"mytest04/crawler/gocrawler/engine"
	"mytest04/crawler/gocrawler/persist"
	"mytest04/crawler/gocrawler/scheduler"
	"mytest04/crawler/gocrawler/zhenai/parser"
	"mytest04/crawler/gocrawler/config"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		//Scheduler: &scheduler.SimpleScheduler{},
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})

	//e.Run(engine.Request{
	//	Url:  "http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc:  parser.ParseCity,
	//})
}

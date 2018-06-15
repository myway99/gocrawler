package main

import (
	"mytest04/crawler/engine"
	"mytest04/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:	"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})

}


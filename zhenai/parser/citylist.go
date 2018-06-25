package parser

import (
	"mytest04/crawler/gocrawler/engine"

	"regexp"
	"mytest04/crawler/gocrawler/config"
)

const cityListRe  = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(
	contents []byte, _ string) engine.ParseResult {

	re := regexp.MustCompile(cityListRe)
	// matchs is a Slice
	matches := re.FindAllSubmatch(contents,-1)
	// [][]string

	result := engine.ParseResult{}
	//limit := 10
	for _, m := range matches {
		//result.Items = append(
		//	result.Items, "City "+ string(m[2]))
		result.Requests= append(
			result.Requests, engine.Request{
				Url: 	string(m[1]),
				Parser: engine.NewFuncParser(
					ParseCity, config.ParseCity),
			})
		//limit--
		//if limit == 0 {
		//	break
		//}
	}

	return result
}
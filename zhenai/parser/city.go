package parser

import (
	"mytest04/crawler/gocrawler/engine"
	"regexp"
	"mytest04/crawler/gocrawler/config"
)

//const cityRe  = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
var (
	profileRe  = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)


func ParseCity(
	contents []byte, _ string) engine.ParseResult {

	//re := regexp.MustCompile(cityRe)
	// matchs is a Slice
	matches := profileRe.FindAllSubmatch(
		contents,-1)
	// [][]string

	result := engine.ParseResult{}
	for _, m := range matches {
		//url := string(m[1])
		//name := string(m[2])
		//result.Items = append(
		//	result.Items, "User "+ name)
		result.Requests= append(
			result.Requests, engine.Request{
				Url: 	string(m[1]),
				// 函数式编程，定义一个匿名函数
				Parser: NewProfileParser(
					string(m[2])),
			})
	}

	matches = cityUrlRe.FindAllSubmatch(
		contents, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests,
			engine.Request{
				Url: string(m[1]),
				Parser:  engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})
	}

	return result
}








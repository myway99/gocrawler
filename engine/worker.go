package engine

import (
	"mytest04/crawler/gocrawler/fetcher"
	"log"
)

func Worker(
	r Request) (ParseResult, error) {

	//log.Printf("fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: eror " +
			"fetching url %s: %v",
			r.Url, err)
		return ParseResult{}, err
	}

	return r.Parser.Parse(body, r.Url), nil
}

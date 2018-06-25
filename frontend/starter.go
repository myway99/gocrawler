package main

import (
	"net/http"
	"mytest04/crawler/gocrawler/frontend/controller"
)

func main() {

	// /css/style.css
	http.Handle("/", http.FileServer(
		http.Dir("crawler/gocrawler/frontend/view")))


	http.Handle("/search",
		controller.CreateSearchResultHandler(
			"crawler/gocrawler/frontend/view/template.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}

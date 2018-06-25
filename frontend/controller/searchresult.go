package controller

import (
	"context"
	"mytest04/crawler/gocrawler/engine"
	"mytest04/crawler/gocrawler/frontend/model"
	"mytest04/crawler/gocrawler/frontend/view"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/olivere/elastic.v5"
	"regexp"
)

// TODO


// support paging
// add start page


type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {

	client, err := elastic.NewClient(
		//ElasticSearch server address and port
		elastic.SetURL("http://192.168.1.188:9200/"),
		// Must turn off sniff in docker
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(
			template),
		client: client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler) ServeHTTP(
	w http.ResponseWriter, req *http.Request) {

	q := strings.TrimSpace(req.FormValue("q"))
	//q = rewriteQueryString(q)

	from, err := strconv.Atoi(
		req.FormValue("from"))
	if err != nil {
		from = 0
	}

	//fmt.Fprintf(w, "q=%s, from=%d", q, from)

	//var page model.SearchResult
	page, err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusBadRequest)
		return
	}
}

const pageSize = 10

func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {

	var result model.SearchResult
	// fill in query string
	result.Query = q

	resp, err := h.client.
		Search("dating_profile").
		Query(elastic.NewQueryStringQuery(
			rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}

	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(
		reflect.TypeOf(engine.Item{}))

	if result.Start == 0 {
		result.PrevFrom = -1
	} else {
		result.PrevFrom =
			(result.Start - 1) /
				pageSize * pageSize
	}

	//result.PrevFrom =
	//	result.Start - len(result.Items)
	//

	result.NextFrom =
		result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")

}
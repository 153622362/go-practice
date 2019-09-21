package controller

import (
	"context"
	vi "go-practice/crawler/frontend/view"
	"reflect"
	"regexp"

	"go-practice/crawler/engine"
	//"fmt"
	"go-practice/crawler/frontend/model"
	"gopkg.in/olivere/elastic.v5"
	"net/http"
	"strconv"
	"strings"
)

//fill in query string
// rewrite query string
//support search button
//support paging
//add start page
type SearchResultHandler struct {
	view   vi.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(
	template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.199.101:9200"),
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	return SearchResultHandler{
		view:   vi.CreateSearchResultView(template),
		client: client,
	}
}

//from分页
// localhost:8888/search?q= ngyhd 本科 &from=0
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	q := strings.TrimSpace(req.FormValue("q"))
	from, err := strconv.Atoi(req.FormValue("from"))

	if err != nil {
		from = 0
	}

	//fmt.Fprint(w, "q=%s, from=%d", q, from)
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

//从elasticsearch中获取数据
func (h SearchResultHandler) getSearchResult(
	q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	result.Query = q
	resp, err := h.client.Search("dating_profile").
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Size(20).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from

	//result.Items = resp.Each(
	//	reflect.TypeOf(engine.Item{}))

	for _, v := range resp.Each(
		reflect.TypeOf(engine.Item{})) {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}

	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)

	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*)`)
	return re.ReplaceAllString(q, "Payload.$1:")
}

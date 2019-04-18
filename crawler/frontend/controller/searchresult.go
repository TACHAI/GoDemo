package controller

import (
	"Demo01/crawler/frontend/view"
	"Demo01/crawler/model"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string)SearchResultHandler  {
	client,err := elastic.NewClient(
		elastic.SetSniff(false))

	if err !=nil{

	}
	return SearchResultHandler{
		view:view.CreateSearchResultView(template),
		client:client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter,req *http.Request) {
	q :=strings.TrimSpace(req.FormValue("q"))
	from,err := strconv.Atoi(req.FormValue("from"))

	if err != nil{
		from =0
	}

	var page model.SearchResult
	page,err := h.getSearchResult(q,from)
	if err !=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	err :=h.view.Render(w,page)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
	fmt.Fprint(w,"q=%s,from=%d",q,from)
}

func (h SearchResultHandler)getSearchResult(
	q string,from int)(model.SearchResult, error)  {
	var result model.SearchResult
	resp,err :=h.client.Search("").
		Query(elastic.NewQueryStringQuery(q)).From(from).Do()
	if err !=nil{
		return result,err
	}
	result.Hits=resp.TotalHits()
	result.Start =from

	return result,err
	//result.Items = resp.Each(reflect.TypeOf())
}
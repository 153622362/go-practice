package view

import (
	"html/template"
	"io"
	"../model"
)

type SearchResultView struct {
	temmlate *template.Template
}

//结果视图
func CreateSearchResultView(
	filename string) SearchResultView {
	return SearchResultView{
		temmlate:template.Must(
			template.ParseFiles(filename)),
	}
}

//渲染视图
func (s SearchResultView) Render (
	w io.Writer, data model.SearchResult) error {
	return s.temmlate.Execute(w, data)
}
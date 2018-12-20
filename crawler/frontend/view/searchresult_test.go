package view

import (
	"../../engine"
	comm "../../model"
	"../model"
	"os"
	"testing"
	vi "../view"
)

func TestSearchResultView_Render(t *testing.T)  {
	view := vi.CreateSearchResultView("template.html")

	out, err := os.Create("template.test.html")
	page := model.SearchResult{}
	//err := template.Execute(os.Stdout, page)
	page.Hits = 10

	item := engine.Item{
		Url : "http://album.zhenai.com/u/1972182204",
		Type: "zhenai",
		Id: "1972182204",
		Payload:comm.Profile{
			Name:		"豆豆",
			Age:		43,
			Height:		161,
			Income:		"5001-8000元",
			Marriage:	"离异",
			Education:	"大学本科",
			Hokou:		"阿坝",
		},
	}

	for i := 0 ; i < 10 ; i++  {
		page.Items = append(page.Items, item)
	}
	err = view.Render(out ,page)
	if err != nil {
		panic(err)
	}
}
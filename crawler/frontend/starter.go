package main

import (
	"go-practice/crawler/frontend/controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(
		`E:\goland\gopath\src\go-practice\crawler\frontend\view`)))

	http.Handle("/search",
		controller.CreateSearchResultHandler(
			`E:\goland\gopath\src\go-practice\crawler\frontend\view\template.html`))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

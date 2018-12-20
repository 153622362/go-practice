package main

import (
	"./controller"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(
		`I:\Google资深工程师深度讲解Go语言\u2pppw\crawler\frontend\view`)))

	http.Handle("/search",
		controller.CreateSearchResultHandler(
			`I:\Google资深工程师深度讲解Go语言\u2pppw\crawler\frontend\view\template.html`))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"net/http"
	"github.com/abnereel/crawler_dating_sites/crawler/frontend/controller"
)

const template = "crawler/frontend/view/template.html"

func main() {
	http.Handle("/", http.FileServer(http.Dir("crawler/frontend/view")))
	http.Handle("/search", controller.CreateSearchResultHandler(template))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}

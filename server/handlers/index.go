package handlers

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed index.html
var indexHtml string

func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, indexHtml)
}

package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Service struct{}

func (serv *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	data := strings.Join(r.Header["Accept"], "")
	var doc string
	switch {
	case data == "*/*":
		w.Header().Add("Content-Type", "application/json")
		doc = `
		{"data": "Hi! This is a json"}
		`
	default:

		w.Header().Add("Content-Type", "text/html")
		doc = fmt.Sprintf(`
<h1>Hello world</h1>
<p>Access Route: %s</p>
	`, r.URL.Path)
	}

	w.Write([]byte(doc))
}

func main() {
	http.ListenAndServe(":3000", &Service{})
}
package main

import (
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	entries map[string]string
}

func (serv *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	var doc string

	if r.Body != nil {
		body, _ := io.ReadAll(r.Body)
		doc = string(body)
	}

	w.Header().Add("Content-Type", "text/plain")

	switch r.Method {
	case http.MethodGet:
		writted, ok := serv.entries[id]

		if !ok {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Write([]byte(writted))
		}
	case http.MethodDelete:
		if _, ok := serv.entries[id]; ok {
			delete(serv.entries, id)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	case http.MethodPost:
		if _, ok := serv.entries[id]; ok {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			serv.entries[id] = doc
		}
	case http.MethodPut:
		if data, ok := serv.entries[id]; ok {
			data += doc
			serv.entries[id] = data
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Not Valid Method:" + r.Method)
	}
}

func main() {
	service := Service{
		entries: map[string]string{},
	}

	http.Handle("/", &service)

	panic(http.ListenAndServe(":3000", nil))
}

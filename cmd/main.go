package main

import (
	"net/http"
)

type handler struct{}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("some text"))
}

func main() {
	var a http.Handler
	a = &handler{}
	_ = http.ListenAndServe("localhost:8080", a)
	http.HandleFunc("/", a.ServeHTTP)

}

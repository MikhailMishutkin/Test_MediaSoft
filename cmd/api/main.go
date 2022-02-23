package main

import (
	"net/http"

	"github.com/MikhailMishutkin/Test_MediaSoft/internal/app/apiserver"
	//"github.com/MikhailMishutkin/Test_MediaSoft/internal/app/apiserver"
)

type handler struct{}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("some text"))
}

func main() {
	config := apiserver.NewConfig()
	s := apiserver.New(config)
	var a http.Handler
	a = &handler{}
	_ = http.ListenAndServe("localhost:8080", a)
	http.HandleFunc("/", a.ServeHTTP)

}

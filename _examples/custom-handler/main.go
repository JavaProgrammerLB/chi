package main

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

//Handler type Handler func;type Mux struct;type Compare interface;
type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func main() {
	// (1)chi.NewRouter()
	r := chi.NewRouter()

	// (2)Mux结构体的Mthod函数
	r.Method("GET", "/", Handler(customHandler))

	// (3)ListenAndServe函数（http包里的函数，两个参数①端口号②Mux）
	http.ListenAndServe(":3333", r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}

package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// package main里的init()函数
func init() {
	chi.RegisterMethod("LINK")
	chi.RegisterMethod("UNLINK")
	chi.RegisterMethod("WOOHOO")
}

// package main里的main()函数
func main() {
	// (1) chi.NewRouter()
	r := chi.NewRouter()

	// Mux结构体的Use函数
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	// Mux结构体的Get函数
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	// Mux结构体的MethodFunc函数
	r.MethodFunc("LINK", "/link", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("custom link method"))
	})
	r.MethodFunc("WOOHOO", "/woo", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("custom woohoo method"))
	})
	r.HandleFunc("/everything", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("capturing all standard http methods, as well as LINK, UNLINK and WOOHOO"))
	})

	// (3)http包里的ListenAndServer函数，第一个参数是端口号，第二个参数是Mux
	http.ListenAndServe(":3333", r)
}

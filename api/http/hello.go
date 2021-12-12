package main

import (
	"log"
	"net/http"
)

func main() {
	//http.HandleFunc("/", hello)
	helloWithMiddleware()

	addr := ":8080"
	log.Println("starting server at", addr)
	_ = http.ListenAndServe(addr, nil)
}

func helloWithMiddleware() {
	m := NewStack()
	m.Use(logTime)
	m.Use(logWrap)
	m.Register("/", http.HandlerFunc(hello))

	for p, h := range m.mux {
		http.Handle(p, h)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello http server \n"))
}

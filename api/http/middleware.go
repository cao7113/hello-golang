package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// https://chai2010.gitbooks.io/advanced-go-programming-book/content/ch5-web/ch5-03-middleware.html

type middleware func(handler http.Handler) http.Handler

type MStack struct {
	stack []middleware
	mux   map[string]http.Handler
}

func NewStack() *MStack {
	return &MStack{mux: make(map[string]http.Handler)}
}

func (m *MStack) Use(w middleware) *MStack {
	m.stack = append(m.stack, w)
	return m
}

func (m *MStack) Register(path string, h http.Handler) {
	mh := h
	for i := len(m.stack) - 1; i >= 0; i-- {
		mh = m.stack[i](mh)
	}
	m.mux[path] = mh
}

var logTime middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t0 := time.Now()
		h.ServeHTTP(w, r)
		d := time.Since(t0)
		log.Infof("taken: %d ms", d.Milliseconds())
	})
}

var logWrap middleware = func(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("begin %+v", r.RequestURI)
		h.ServeHTTP(w, r)
		<-time.After(35 * time.Millisecond)
		log.Infof("end %+v", r.RequestURI)
	})
}

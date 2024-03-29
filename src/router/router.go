package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

type Route struct {
	pattern string
	handler http.Handler
	*mux.Route
}

type GuruRouter struct {
	routes []*Route
}

func NewRouter() {
	router = mux.NewRouter()
}

func (h *GuruRouter) Handler(method string, pattern string, handler http.Handler) {
	router.Handle(pattern, handler).Methods(method)
}

func (h *GuruRouter) HandleFunc(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(pattern, handler).Methods(method)
}

func (h *GuruRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func ExtractUrlWildcard(r *http.Request, param string) string {
	params := mux.Vars(r)
	return params[param]
}

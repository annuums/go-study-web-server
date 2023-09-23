package app

import (
	"net/http"

	"github.com/annuums/go-study-web-server/handlers"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/home/", http.StripPrefix("/home", &handlers.HomeRouter{}))
	mux.Handle("/", http.NotFoundHandler())

	return mux
}

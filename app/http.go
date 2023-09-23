package app

import (
	"net/http"

	routers "github.com/annuums/go-study-web-server/routes"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/home/", http.StripPrefix("/home", &routers.HomeRouter{}))
	mux.Handle("/", http.NotFoundHandler())

	return mux
}

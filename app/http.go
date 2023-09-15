package app

import (
	"net/http"

	"github.com/annuums/go-study-web-server/handlers"
)

func NewHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", handlers.HomeHandler)

	return mux
}
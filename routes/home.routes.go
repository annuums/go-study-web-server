package routers

import (
	"fmt"
	"net/http"

	"github.com/annuums/go-study-web-server/handlers"
)

type HomeRouter struct{}

/*
*

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
func (home *HomeRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Requested %s\n", req.URL.Path)

	handler := &handlers.HomeHandler{}
	switch req.URL.Path {
	case "/":
		handler.Handles(res, req)
	default:
		http.NotFound(res, req)
	}
}

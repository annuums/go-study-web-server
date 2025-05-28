package routers

import (
	"fmt"
	"net/http"

	"github.com/annuums/go-study-web-server/handlers"
)

type StudentRouter struct{}

/*
*

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}
*/
func (home *StudentRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Requested %s\n", req.URL.Path)

	handler := &handlers.StudentHandler{}

	switch req.URL.Path {
	case "/:id":
		handler.Handles(res, req)
	case "/":
		handler.Handles(res, req)
	default:
		http.NotFound(res, req)
	}
}

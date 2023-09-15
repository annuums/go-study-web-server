package handlers

import (
	"fmt"
	"net/http"
)

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprintln(res, "Hello, Home!")
}
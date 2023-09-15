package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct {}

/**
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (home *HomeHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Requested %s\n", req.URL.Path)
	
	switch req.URL.Path {
	case "/":
		indexHandler(res, req)
	case "/test":
		testHandler(res, req)
	default:
		http.NotFound(res, req)
	}
}

func indexHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getIndex(res, req)
	case http.MethodPost:
		postIndex(res, req)
	}
}

func getIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home")
}

func postIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Helo, This is Post Handler! You can [GET, POST] to /home")
}

func testHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		getTest(res, req)
	case http.MethodPost:
		postTest(res, req)
	}	
}

func getTest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home/test")
}

func postTest(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Post Handler! You can [GET, POST] to /home/test")
}
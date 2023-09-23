package handlers

import (
	"fmt"
	"net/http"
)

type HomeHandler struct{}

func (handler *HomeHandler) Handles(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		handler.getIndex(res, req)
	case http.MethodPost:
		handler.postIndex(res, req)
	}
}

func (handler *HomeHandler) getIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Hello, This is Get Handler! You can [GET, POST] to /home")
}

func (handler *HomeHandler) postIndex(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Helo, This is Post Handler! You can [GET, POST] to /home")
}

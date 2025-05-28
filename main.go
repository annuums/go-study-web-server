package main

import (
	"log"
	"net/http"

	"github.com/annuums/go-study-web-server/app"
)

func main() {
	log.Println("Server is running on :3000...")
	http.ListenAndServe(":3000", app.NewHandler())
}
package main

import (
	"log"
	"net/http"

	"github.com/annuums/go-study-web-server/app"
)

func main() {

	log.Println("Server is running on :3000...")
	err := http.ListenAndServe(":3000", app.NewHandler())
	if err != nil {

		log.Fatalf("Failed to start server: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"myself/internal/api"
	"myself/pkg/data"
	"net/http"
)

func main() {
	fmt.Println("🚀 Starting MySelf API on http://localhost:8080")

	// Store = accès aux données
	store := data.NewStore("records")

	// Handlers = logique API
	handlers := api.Handlers{
		Store: store,
	}

	// Router = routes + middlewares
	router := api.NewRouter(handlers)

	// Serveur HTTP
	log.Fatal(http.ListenAndServe(":8080", router))
}

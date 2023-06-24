package main

import (
	"gomicroservice/api"
	"gomicroservice/repository"
	"gomicroservice/service"
	"log"
	"net/http"
)

func main() {
	// Initialize new repository
	repo := repository.NewRepository()

	// Initialize new service with the repository as a dependency
	serv := service.NewService(repo)

	// Initialize new handler with the service as a dependency
	handler := api.NewHandler(serv)

	// Get the routes from the handler
	router := handler.GetRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}

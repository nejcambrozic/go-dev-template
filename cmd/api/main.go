package main

import (
	"github.com/nejcambrozic/go-dev-tempalte/pkg/health"
	"github.com/nejcambrozic/go-dev-tempalte/pkg/http/rest"
	"log"
	"net/http"
)

func main() {

	// init all services needed to handle requests
	healthChecker := health.NewService()

	// pass all the services into a handler
	router := rest.Handler(healthChecker)

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":3000", router))
}

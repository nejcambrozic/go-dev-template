package rest

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/nejcambrozic/go-dev-tempalte/pkg/health"
	"net/http"
)

// Setups the routes and handler functions
// each handler function gets passed the service interface so it can call the needed method
func Handler(h health.Service) http.Handler {
	router := httprouter.New()

	// This is how you define a handler that uses a service
	router.GET("/health", getHealth(h))
	// Can also do it like this for simple functions
	router.GET("/ping", pingResponse)

	return router
}

func getHealth(h health.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Call the service implementation to get health information
		healthMessage, status := h.GetHealth()

		// Always do this: set return type to json
		w.Header().Set("Content-Type", "application/json")
		// Return response status and message returned by the service
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(healthMessage)
	}
}

func pingResponse(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("pong")
}

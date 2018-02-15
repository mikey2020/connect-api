package router

import (
	"github.com/gorilla/mux"
	. "connect/controller"
	. "connect/middlewares"
)

func SetConceptRoutes(conceptRouter *mux.Router) *mux.Router {
	// baseRouter := mux.NewRouter()
	// conceptRouter := baseRouter.PathPrefix("/api/v1").Subrouter()
	conceptRouter.HandleFunc("/api/v1/concept", LoggingMiddleware(Authenticate(AddConcept))).Methods("POST")
	conceptRouter.HandleFunc("/api/v1/user/concepts",LoggingMiddleware(Authenticate(GetUserConcepts))).Methods("GET")

	return conceptRouter
}
package router

import (
	"github.com/gorilla/mux"
)

// InitRoutes function for initialising routes
func InitRoutes() *mux.Router {

	router := mux.NewRouter().StrictSlash(false)

	// Index route
	router = setIndexRoute(router)

	//Routes for concept entity
	router = setConceptRoutes(router)

	// Routes for the User entity
	router = setUserRoutes(router)

	return router
}

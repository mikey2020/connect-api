package router

import (
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router{

	router := mux.NewRouter().StrictSlash(false)

	// Index route
	router = SetIndexRoute(router)
	
	//Routes for concept entity
	router = SetConceptRoutes(router)

	// Routes for the User entity
	router = SetUserRoutes(router)
	
	return router
}


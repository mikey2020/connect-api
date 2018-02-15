package router

import (
	"github.com/gorilla/mux"
	. "connect/controller"
	. "connect/middlewares"
)


func SetIndexRoute(indexRouter *mux.Router) *mux.Router{
	// indexRouter := mux.NewRouter()

	indexRouter.HandleFunc("/", LoggingMiddleware(Index)).Methods("GET")

    return indexRouter

}
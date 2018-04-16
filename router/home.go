package router

import (
	mwr "github.com/mikey2020/connect-api/middlewares"

	ctrl "github.com/mikey2020/connect-api/controller"

	"github.com/gorilla/mux"
)

func setIndexRoute(indexRouter *mux.Router) *mux.Router {
	// indexRouter := mux.NewRouter()

	indexRouter.HandleFunc("/", mwr.LoggingMiddleware(ctrl.Index)).Methods("GET")

	return indexRouter

}

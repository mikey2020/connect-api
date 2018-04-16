package router

import (
	mwr "github.com/mikey2020/connect-api/middlewares"

	ctrl "github.com/mikey2020/connect-api/controller"

	"github.com/gorilla/mux"
)

func setConceptRoutes(conceptRouter *mux.Router) *mux.Router {
	// baseRouter := mux.NewRouter()
	// conceptRouter := baseRouter.PathPrefix("/api/v1").Subrouter()
	conceptRouter.HandleFunc("/api/v1/concept", mwr.LoggingMiddleware(mwr.Authenticate(ctrl.AddConcept))).Methods("POST")
	conceptRouter.HandleFunc("/api/v1/user/concepts", mwr.LoggingMiddleware(mwr.Authenticate(ctrl.GetUserConcepts))).Methods("GET")
	conceptRouter.HandleFunc("/api/v1/{concept_id}/update", mwr.LoggingMiddleware(mwr.Authenticate(ctrl.EditConcept))).Methods("PUT")
	conceptRouter.HandleFunc("/api/v1/{concept_id}/join", mwr.LoggingMiddleware(mwr.Authenticate(ctrl.JoinConcept))).Methods("POST")
	return conceptRouter
}

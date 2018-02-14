package router

import (
    "github.com/gorilla/mux"
	. "connect/controller"
	. "connect/middlewares"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	baseRouter := mux.NewRouter()
	userRouter := baseRouter.PathPrefix("/api/v1").Subrouter()
	userRouter.HandleFunc("/user/signup", LoggingMiddleware(CreateUser)).Methods("POST")
	userRouter.HandleFunc("/users", LoggingMiddleware(Authenticate(GetAllUsers))).Methods("GET")
	userRouter.HandleFunc("/user/signin", LoggingMiddleware(SignInUser)).Methods("POST")
    return userRouter
}
package router

import (
    "github.com/gorilla/mux"
	. "connect/controller"
	. "connect/middlewares"
)

func SetUserRoutes(userRouter *mux.Router) *mux.Router {
	// baseRouter := mux.NewRouter()
	userRouter.HandleFunc("/api/v1/user/signup", LoggingMiddleware(CreateUser)).Methods("POST")
	userRouter.HandleFunc("/api/v1/users", LoggingMiddleware(Authenticate(GetAllUsers))).Methods("GET")
	userRouter.HandleFunc("/api/v1/user/signin", LoggingMiddleware(SignInUser)).Methods("POST")
    return userRouter
}
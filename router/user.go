package router

import (
	mwr "github.com/mikey2020/connect-api/middlewares"

	ctrl "github.com/mikey2020/connect-api/controller"

	"github.com/gorilla/mux"
)

func setUserRoutes(userRouter *mux.Router) *mux.Router {
	// baseRouter := mux.NewRouter()
	userRouter.HandleFunc("/api/v1/user/signup", mwr.LoggingMiddleware(ctrl.CreateUser)).Methods("POST")
	userRouter.HandleFunc("/api/v1/users", mwr.LoggingMiddleware(mwr.Authenticate(ctrl.GetAllUsers))).Methods("GET")
	userRouter.HandleFunc("/api/v1/user/signin", mwr.LoggingMiddleware(ctrl.SignInUser)).Methods("POST")
	return userRouter
}

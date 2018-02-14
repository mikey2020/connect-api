package controller

import (
	"net/http"
	help "connect/helper"
)

type Controller struct {}

func Index(w http.ResponseWriter, r *http.Request){
	help.RespondWithJson(w, 200, "Welcome to connect Api")
}
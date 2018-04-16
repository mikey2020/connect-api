package controller

import (
	"net/http"

	. "github.com/mikey2020/connect-api/mongo"

	help "github.com/mikey2020/connect-api/helper"
)

var Dao = DAO{}

type Controller struct{}

func Index(w http.ResponseWriter, r *http.Request) {
	help.RespondWithJson(w, 200, "Welcome to connect Api")
}

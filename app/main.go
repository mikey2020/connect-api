package app

import (
	r "github.com/mikey2020/connect-api/router"

	ctrl "github.com/mikey2020/connect-api/controller"

	"github.com/gorilla/mux"
)

//Router  - to initialize all routes
var Router *mux.Router

// SetupConfig - set up database config
func SetupConfig(server string, database string) {
	ctrl.Dao.Server = server
	ctrl.Dao.Database = database
}

// Init to help initialize database
func Init() {
	Router = r.InitRoutes()
	ctrl.Dao.Connect()
}

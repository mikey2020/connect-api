package main

import (
	"os"
	"log"
	"net/http"
	"github.com/subosito/gotenv"
	. "connect/mongo"
	. "connect/router"
)

var dao = DAO{}

func init() {
	gotenv.Load()
	dao.Server = os.Getenv("Server")
	dao.Database = os.Getenv("Database")
	dao.Connect()
}

func main(){
	router := InitRoutes()
	log.Println("Serving listening on port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
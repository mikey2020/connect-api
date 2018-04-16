package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mikey2020/connect-api/app"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	app.SetupConfig(os.Getenv("SERVER"), os.Getenv("DATABASE"))
	app.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("Serving listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}

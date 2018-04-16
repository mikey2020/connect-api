package tests

import (
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	ctrl "github.com/connect-api/controller"

	"github.com/connect-api/app"

	model "github.com/connect-api/mongo"

	. "github.com/haoxins/supertest"
	"github.com/subosito/gotenv"
)

func TestMain(m *testing.M) {
	gotenv.Load()
	app.SetupConfig(os.Getenv("SERVER"), os.Getenv("TEST_DATABASE"))
	fmt.Println(os.Getenv("TEST_DATABASE"))
	app.Init()
	os.Exit(m.Run())
}

func TestUser(t *testing.T) {
	testServer := httptest.NewServer(app.Router)
	defer testServer.Close()

	t.Run("Can Sign up", func(t *testing.T) {
		Request(testServer.URL, t).
			Post("/api/v1/user/signup").
			Send(model.User{Username: "test", Email: "test@gmail.com", Password: "testing"}).
			Expect(201).
			Expect("Content-Type", "application/json").
			End()
	})

	t.Run("Can Sign in", func(t *testing.T) {
		Request(testServer.URL, t).
			Post("/api/v1/user/signin").
			Send(model.User{Username: "test", Password: "testing"}).
			Expect(200).
			Expect("Content-Type", "application/json").
			End()
	})

	error := ctrl.Dao.RemoveAllUsers()

	if error != nil {
		panic(error)
	}

}

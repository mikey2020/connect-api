package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	. "github.com/mikey2020/connect-api/mongo"

	"github.com/dgrijalva/jwt-go"
	"github.com/subosito/gotenv"
	gomail "gopkg.in/gomail.v2"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJson(w, code, map[string]string{"error": msg})
}

func RespondWithJsonError(w http.ResponseWriter, code int, payload interface{}) {
	RespondWithJson(w, code, map[string]interface{}{"error": payload})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func RespondWithSuccess(w http.ResponseWriter, code int, msg string, key string) {
	RespondWithJson(w, code, map[string]string{key: msg})
}

func GenerateToken(w http.ResponseWriter, r *http.Request, user User) (string, error) {
	gotenv.Load()
	_ = json.NewDecoder(r.Body).Decode(&user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":      user.ID,
		"username": user.Username,
		"password": user.Password,
		"email":    user.Email,
	})
	tokenString, error := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return tokenString, error
}

// Send Mail to users
func SendMail(receipents []string, subject, message string) {
	gotenv.Load()
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("USER_EMAIL"))
	for _, person := range receipents {
		m.SetHeader("To", person)
	}
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", message)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("USER_EMAIL"), os.Getenv("PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
}

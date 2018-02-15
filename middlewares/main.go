package middlewares

import (
	"net/http"
	. "connect/helper"
	"log"
	"time"
	"strings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"os"
	"github.com/subosito/gotenv" 
	"fmt"
)

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
    })
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	gotenv.Load()
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authorizationHeader := r.Header.Get("authorization")
        if authorizationHeader != "" {
            bearerToken := strings.Split(authorizationHeader, " ")
            if len(bearerToken) == 2 {
                token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, fmt.Errorf("There was an error")
                    }
                    return []byte(os.Getenv("JWT_SECRET")), nil
                })
                if error != nil {
                    RespondWithError(w,401,error.Error())
                    return
                }
                if token.Valid {
                    context.Set(r, "decoded", token.Claims)
                    next(w, r)
                } else {
                    RespondWithError(w,401,"Invalid authorization token")
                }
            }
        } else {
            RespondWithError(w,401,"An authorization header is required")
        }
    })
}
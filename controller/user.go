package controller

import (
	"encoding/json"
	"net/http"

	. "github.com/mikey2020/connect-api/validations"

	. "github.com/mikey2020/connect-api/mongo"

	. "github.com/mikey2020/connect-api/helper"

	"gopkg.in/mgo.v2/bson"
)

// CreateUser -
func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	} else {
		_, userErr := Dao.FindUser("username", user.Username)
		_, emailErr := Dao.FindUser("email", user.Email)
		if userErr == nil || emailErr == nil {
			RespondWithError(w, 409, "User already exists")
			return
		} else if err, errStatus := ValidateSignUpRequest(user); errStatus != true {
			user.ID = bson.NewObjectId()
			user.Password, _ = user.HashPassword(user.Password)
			if err := Dao.InsertUser(user); err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			receipent := []string{user.Email}
			go SendMail(receipent, "Welcome", "Welcome to <h2>Connect</h2>")
			token, err := GenerateToken(w, r, user)
			if err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			RespondWithSuccess(w, 201, token, "userToken")
		} else {
			RespondWithJsonError(w, 400, err)
		}

	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := Dao.FindAllUsers()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, users)
}

func SignInUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
	} else {
		foundUser, userErr := Dao.FindUser("username", user.Username)
		// _, emailErr := Dao.FindUser("email", user.Email)
		if userErr != nil {
			RespondWithError(w, 401, "Invalid User")
		} else {
			if user.CheckPasswordHash(foundUser.Password, user.Password) == true {
				token, err := GenerateToken(w, r, foundUser)
				if err == nil && token != "" {
					RespondWithSuccess(w, 200, token, "userToken")
				}
				return
			}
			RespondWithError(w, 401, "Invalid Sign in parameters")
		}
	}

}

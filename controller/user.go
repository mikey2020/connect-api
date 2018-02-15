package controller

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"encoding/json"
	. "connect/mongo"
	. "connect/helper"
	. "connect/validations"
)

var dao = DAO{}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	} else {
		_, userErr := dao.FindUser("username",user.Username)
		_, emailErr := dao.FindUser("email",user.Email)
		if userErr  == nil || emailErr == nil {
			RespondWithError(w, 409, "User already exists")
			return 
		} else if err,errStatus := ValidateSignUpRequest(user); errStatus != true {
			user.ID = bson.NewObjectId()
			user.Password, _ = user.HashPassword(user.Password)
			if err := dao.InsertUser(user); err != nil {
				RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			}
			receipent := []string{user.Email}
			SendMail(receipent,"Welcome","Welcome to <h2>Connect</h2>")
			GenerateToken(w, r, user)
		} else {
			RespondWithJsonError(w,400,err)
		}

	}
}

func GetAllUsers(w http.ResponseWriter, r *http.Request){
	users, err := dao.FindAllUsers()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, users)
}

func SignInUser(w http.ResponseWriter, r *http.Request){
	defer r.Body.Close()
	var user User
	
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	} else {
		foundUser, userErr := dao.FindUser("username", user.Username)
		_, emailErr := dao.FindUser("email", user.Email)
		if userErr != nil || emailErr != nil {
			RespondWithError(w, 401, "Invalid Sign in parameters")
		} else {
			if user.CheckPasswordHash(foundUser.Password, user.Password) == true {
				GenerateToken(w, r, foundUser)
				return
			} 
			RespondWithError(w, 401, "Invalid sign in parameters")
		}
	}

}
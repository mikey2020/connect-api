package validations

import (
	"regexp"
	"strings"

	. "github.com/mikey2020/connect-api/mongo"
)

const EXP_EMAIL = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

var fallback interface{}

func CreateErrorMessage(key string, value string) interface{} {
	return map[string]string{key: value}
}

func ValidateSignUpRequest(user User) (interface{}, bool) {

	re := regexp.MustCompile(EXP_EMAIL)

	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if user.Username == "" {
		return CreateErrorMessage("username", "Username is required"), true
	}
	if user.Password == "" {
		return CreateErrorMessage("password", "Password is required"), true
	}
	if user.Email == "" {
		return CreateErrorMessage("email", "Email is required"), true
	} else if re.MatchString(user.Email) != true {
		return CreateErrorMessage("email", "Please enter a valid email"), true
	}

	return fallback, false
}

func ValidateSignInRequest(user User) (interface{}, bool) {
	re := regexp.MustCompile(EXP_EMAIL)

	user.Email = strings.TrimSpace(user.Email)

	if user.Email == "" {
		return CreateErrorMessage("email", "Email is required"), true
	} else if re.MatchString(user.Email) != true {
		return CreateErrorMessage("email", "Please enter a valid email"), true
	}
	if user.Password == "" {
		return CreateErrorMessage("password", "Password is required"), true
	}

	return fallback, false
}

func ValidateConcept(concept Concept) (interface{}, bool) {
	concept.Topic = strings.TrimSpace(concept.Topic)
	concept.Description = strings.TrimSpace(concept.Description)

	if concept.Topic == "" {
		return CreateErrorMessage("topic", "Topic is required"), true
	}
	if concept.Description == "" {
		return CreateErrorMessage("description", "Description is required"), true
	}

	return fallback, false
}

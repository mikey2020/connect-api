package controller

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	. "connect/mongo"
	. "connect/helper"
	. "connect/validations"
	"github.com/gorilla/context"
)


func AddConcept(w http.ResponseWriter,r *http.Request){
	defer r.Body.Close()
    var concept Concept
	 
	if err := json.NewDecoder(r.Body).Decode(&concept); err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return 
	} else {
		decoded := context.Get(r, "decoded")
		id := decoded.(jwt.MapClaims)["_id"]
		_,conErr := dao.FindConcept("topic", concept.Topic)
		if conErr == nil {
			RespondWithError(w, 409, "Concept already exists")
			return
		} else {
			if err, errStatus := ValidateConcept(concept); !errStatus {
				concept.ID = bson.NewObjectId()
				concept.UserId = id
				if err := dao.InsertConcept(concept); err != nil {
					RespondWithError(w, http.StatusInternalServerError, err.Error())
					return
				}
				RespondWithSuccess(w, 201, fmt.Sprintf("%s concept created", concept.Topic), "message")
			} else {
                RespondWithJsonError(w, 400, err)
			}
		}
	}
}

func GetUserConcepts(w http.ResponseWriter, r *http.Request){
	decoded := context.Get(r, "decoded")
	id := decoded.(jwt.MapClaims)["_id"]
	concepts, err := dao.FindConceptsById(id)
	fmt.Println(id.(string));
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, concepts)
}


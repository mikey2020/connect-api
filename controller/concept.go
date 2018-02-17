package controller

import (
	"time"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"encoding/json"
	. "connect/mongo"
	. "connect/helper"
	. "connect/validations"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
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
				concept.Users = append(concept.Users, decoded.(jwt.MapClaims)["email"].(string))
				concept.CreatedAt = time.Now()
				concept.UpdatedAt = time.Now()
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
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, concepts)
}

func EditConcept(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	var concept Concept
	if err := json.NewDecoder(r.Body).Decode(&concept); err != nil {
		RespondWithError(w, 401, "Invalid Request Payload")
	}
	savedConcept, err := dao.FindConceptById(string(vars["concept_id"]))
	if err != nil {
		RespondWithError(w, 404, "Concept does not exist")
		return
	}
	concept.ID = savedConcept.ID
	concept.UpdatedAt = time.Now()
	if err := dao.UpdateConcept(concept, savedConcept.ID); err != nil {
		fmt.Println(err)
		RespondWithError(w, 500, "Internal Server Error")
	} else {
		RespondWithSuccess(w, 200, "Concept updated", "message")
	} 

}

func JoinConcept(w http.ResponseWriter, r *http.Request) {
	 params := mux.Vars(r)
	 decoded := context.Get(r, "decoded")
	 con,err := dao.FindConceptById(params["concept_id"])
	 if err != nil {
		 RespondWithError(w, 404, "Concept does not exist")
	 } else {
		con.Users = append(con.Users, decoded.(jwt.MapClaims)["email"].(string))
		if err := dao.UpdateConcept(con, con.ID); err != nil {
			RespondWithError(w, 500, "Internal Server Error")
		} else {
			RespondWithSuccess(w, 200, "joined concept successfully", "message")
			return
		} 
	 }
}
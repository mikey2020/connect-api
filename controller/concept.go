package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	. "github.com/mikey2020/connect-api/validations"

	. "github.com/mikey2020/connect-api/mongo"

	. "github.com/mikey2020/connect-api/helper"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func AddConcept(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var concept Concept

	if err := json.NewDecoder(r.Body).Decode(&concept); err != nil {
		RespondWithError(w, 400, "Invalid request payload")
		return
	} else {
		decoded := context.Get(r, "decoded")
		id := decoded.(jwt.MapClaims)["_id"]
		_, conErr := Dao.FindConcept("topic", concept.Topic)
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
				if err := Dao.InsertConcept(concept); err != nil {
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

func GetUserConcepts(w http.ResponseWriter, r *http.Request) {
	decoded := context.Get(r, "decoded")
	id := decoded.(jwt.MapClaims)["_id"]
	concepts, err := Dao.FindConceptsById(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	RespondWithJson(w, http.StatusOK, concepts)
}

func EditConcept(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var concept Concept
	if err := json.NewDecoder(r.Body).Decode(&concept); err != nil {
		RespondWithError(w, 401, "Invalid Request Payload")
	}
	savedConcept, err := Dao.FindConceptById(string(vars["concept_id"]))
	if err != nil {
		RespondWithError(w, 404, "Concept does not exist")
		return
	}
	concept.ID = savedConcept.ID
	concept.UpdatedAt = time.Now()
	if err := Dao.UpdateConcept(concept, savedConcept.ID); err != nil {
		fmt.Println(err)
		RespondWithError(w, 500, "Internal Server Error")
	} else {
		RespondWithSuccess(w, 200, "Concept updated", "message")
	}

}

func JoinConcept(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	decoded := context.Get(r, "decoded")
	con, err := Dao.FindConceptById(params["concept_id"])
	if err != nil {
		RespondWithError(w, 404, "Concept does not exist")
	} else {
		con.Users = append(con.Users, decoded.(jwt.MapClaims)["email"].(string))
		if err := Dao.UpdateConcept(con, con.ID); err != nil {
			RespondWithError(w, 500, "Internal Server Error")
		} else {
			RespondWithSuccess(w, 200, "joined concept successfully", "message")
			return
		}
	}
}

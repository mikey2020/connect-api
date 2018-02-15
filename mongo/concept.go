package mongo

import (
	"gopkg.in/mgo.v2/bson"
)

const CONCEPT_COLLECTION = "Concepts"


// get all concepts in database
func (m *DAO) FindAllConcepts() ([]Concept, error) {
	var concepts []Concept
	err := db.C(CONCEPT_COLLECTION).Find(bson.M{}).All(&concepts)
	return concepts, err
}

// get some concepts in database
func (m *DAO) FindConceptsById(id interface{}) ([]Concept, error) {
	var concepts []Concept
	err := db.C(CONCEPT_COLLECTION).Find(bson.M{"user_id": id }).All(&concepts)
	return concepts, err
}

// get a concept by Id
func (m *DAO) FindConceptById(id string) (Concept, error) {
	var concept Concept
	err := db.C(CONCEPT_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&concept)
	return concept, err
}

// get concept by any of its atrributes
func (m *DAO) FindConcept(field, value string) (Concept,error) {
	var concept Concept
	err := db.C(CONCEPT_COLLECTION).Find(bson.M{field: value }).One(&concept)

	return concept,err
}

// add a new concept to the database
func (m *DAO) InsertConcept(concept Concept) error {
	err := db.C(CONCEPT_COLLECTION).Insert(&concept)
	return err
}

// remove a concept from the database
func (m *DAO) DeleteConcept(concept Concept) error {
	err := db.C(CONCEPT_COLLECTION).Remove(&concept)
	return err
}

// update a concept in the database
func (m *DAO) UpdateConcept(concept Concept) error {
	err := db.C(CONCEPT_COLLECTION).UpdateId(concept.ID, &concept)
	return err
}
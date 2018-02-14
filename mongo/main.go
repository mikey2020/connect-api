package mongo


import (
	"log"
	mgo "gopkg.in/mgo.v2"
)

type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database


func (m *DAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}



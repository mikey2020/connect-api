package mongo

import "gopkg.in/mgo.v2/bson"

// define user collection
const (
	COLLECTION = "Users"
)

// get all users in database
func (m *DAO) FindAll() ([]User, error) {
	var users []User
	err := db.C(COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// get a user by Id
func (m *DAO) FindById(id string) (User, error) {
	var user User
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// get user by any of its atrributes
func (m *DAO) Find(field, value string) (User,error) {
	var user User
	err := db.C(COLLECTION).Find(bson.M{field: value }).One(&user)

	return user,err
}

// add a new user to the database
func (m *DAO) Insert(user User) error {
	err := db.C(COLLECTION).Insert(&user)
	return err
}

// remove a user from the database
func (m *DAO) Delete(user User) error {
	err := db.C(COLLECTION).Remove(&user)
	return err
}

// update a user in the database
func (m *DAO) Update(user User) error {
	err := db.C(COLLECTION).UpdateId(user.ID, &user)
	return err
}
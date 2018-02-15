package mongo

import "gopkg.in/mgo.v2/bson"

// define user collection
const USER_COLLECTION = "Users"


// get all users in database
func (m *DAO) FindAllUsers() ([]User, error) {
	var users []User
	err := db.C(USER_COLLECTION).Find(bson.M{}).All(&users)
	return users, err
}

// get a user by Id
func (m *DAO) FindUserById(id string) (User, error) {
	var user User
	err := db.C(USER_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// get user by any of its atrributes
func (m *DAO) FindUser(field, value string) (User,error) {
	var user User
	err := db.C(USER_COLLECTION).Find(bson.M{field: value }).One(&user)

	return user,err
}

// add a new user to the database
func (m *DAO) InsertUser(user User) error {
	err := db.C(USER_COLLECTION).Insert(&user)
	return err
}

// remove a user from the database
func (m *DAO) DeleteUser(user User) error {
	err := db.C(USER_COLLECTION).Remove(&user)
	return err
}

// update a user in the database
func (m *DAO) Update(user User) error {
	err := db.C(USER_COLLECTION).UpdateId(user.ID, &user)
	return err
}
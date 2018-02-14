package mongo


import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Email       string        `bson:"email" json:"email"`
	Password    string        `bson:"password" json:"password"`
}

func (u User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u User) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
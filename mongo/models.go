package mongo


import (
	"gopkg.in/mgo.v2/bson"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Username    string        `bson:"username" json:"username"`
	Email       string        `bson:"email" json:"email"`
	Password    string        `bson:"password" json:"password"`
	Concepts    []Concept
}

func (u User) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (u User) CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type Concept struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	UserId      interface{} `bson:"user_id,omitempty" json:"user_id"`
	Topic       string        `bson:"topic,omitempty" json:"topic"`
	Description string        `bson:"description,omitempty" json:"description"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
}


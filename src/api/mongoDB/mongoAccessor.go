package mongoDB

import (
	"api/auth"
	"context"
	"fmt"
	"log"
	// "reflect"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type MongoAccessor struct {
	Users *mongo.Collection
}

func NewUserAccess() auth.UserAccess {
	return &MongoAccessor{
		Users: Users,
	}
}

//SignUp gets user struct, saves it in DB, returns error
func (m *MongoAccessor) SignUp(user auth.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(user.HashedPassword, bcrypt.DefaultCost)
	if err != nil {
		return auth.ErrorMessage
	}
	user.HashedPassword = hashedPassword
		fmt.Println(user)

	// fmt.Println(m.Users)
	insertResult, err := m.Users.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Errorf("SignUp Complete! Welcome %s!", insertResult.InsertedID)
}

//Login gets userName and password,
func (m *MongoAccessor) LogIn(username string, password auth.PASSWORD) (auth.User, error) {
	filter := bson.D{{"username", username}}
	cur, err := m.Users.Find(context.TODO(), filter)
	if err != nil {
		return auth.User{}, auth.ErrorMessage
	}
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem auth.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		// err := bcrypt.CompareHashAndPassword()
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return auth.User{}, nil
}

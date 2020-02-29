//functions in this file controls the data in database.

package auth

import (
	"errors"
	"time"
	
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Casting Type
type PASSWORD []byte
//primitive.ObjectID is a type of mongoDB ID 
// type ID primitive.ObjectID

//Error
var ErrorMessage = errors.New("No user found")

type User struct {
	Id			   primitive.ObjectID  `json:"id" bson:"_id"`
	Username       string    `json:"username"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName`
	HashedPassword PASSWORD  `json:"hashedPassword"`
	SignedUpDate   time.Time `json:"signedUpDate,omitempty"`
}

//slice of user structs
var users = []User{
	{
		Username:       "hko",
		FirstName:      "Hyeonjung",
		LastName:       "Ko",
		HashedPassword: PASSWORD("asdfjsdkl123a"),
		SignedUpDate:   time.Now(),
	},
}

//userAccess is an Interface for user database requires
type UserAccess interface {
	SignUp(u User) (string, error)                                   //signup
	LogIn(userName string, password PASSWORD) (string, error) //login
	// LogOut() error                          //logout
	// Check(Id ID) (user, error)              //check if user is loggedin
}




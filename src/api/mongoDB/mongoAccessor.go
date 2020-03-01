package mongoDB

import (
	"api/auth"
	"context"
	"fmt"
	"log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

//.Collection("users")


func DB() *mongo.Database {
	return GetClient().Database("PlanA")
}

type UserAccessor struct {
	Users *mongo.Collection
}

//userAccess interface (login, signup)
func NewUserAccess() auth.UserAccess {
	return &UserAccessor{
		Users: DB().Collection("users"),
	}
}

// func NewLocationAccess() auth.LocationAccess {
// 	return &MongoAccessor{
// 		Location : DB().Collection("location")
// 	}
// }

//Change it to UserAccessor?

//SignUp gets user struct, saves it in DB, returns error
func (m *UserAccessor) SignUp(user auth.User) (string, error) {
	//check if there is matching username
	var elem auth.User
	filter := bson.D{{"username", user.Username}}
	errFind := m.Users.FindOne(context.TODO(), filter).Decode(&elem)
	if errFind != mongo.ErrNoDocuments {
		return "", fmt.Errorf("choose different username!")
	}
	//create new hashedPassword
	hashedPassword, errHash := bcrypt.GenerateFromPassword(user.HashedPassword, bcrypt.DefaultCost)
	if errHash != nil {
		return "", auth.ErrorMessage
	}
	user.HashedPassword = hashedPassword
	//insert user in DB
	insertResult, errInsert := m.Users.InsertOne(context.TODO(), user)
	if errInsert != nil {
		log.Fatal(errInsert)
	}
	//convert ID to hex
	userID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	fmt.Println(userID)
	return userID, nil
}

//Login gets userName and password,
func (m *UserAccessor) LogIn(username string, password auth.PASSWORD) (string, error) {
	filter := bson.D{{"username", username}}
	var elem auth.User
	err := m.Users.FindOne(context.TODO(), filter).Decode(&elem)
	if err != nil {
	//IT DOES NOT RETURN ID... 
    // ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			return "", auth.ErrorMessage
		}
		log.Fatal(err)
	}
	// fmt.Println(reflect.TypeOf(elem))
	fmt.Println("Found user", &elem)
	fmt.Println(elem)
	userID := elem.Id.Hex()
	fmt.Println(userID)
	return userID, nil
}

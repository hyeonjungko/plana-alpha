package mongoDB

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"

)



// var Users *mongo.Collection = new(mongo.Collection)
var Users

func StartDB() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reflect.TypeOf(client))
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	// DB = client.Database("PlanA")
	Users = client.Database("PlanA").Collection("users")
	// Collections = createCollection(client);
	// Users.InsertOne(context.TODO(), bson.M{"name": "pi", "value": 3.14159})
	fmt.Println("Connected to MongoDB!")
}




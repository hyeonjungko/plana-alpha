package mongoDB

import (
	"context"
	"fmt"
	"log"
	"os"
		
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/joho/godotenv"
)



// var Users *mongo.Collection = new(mongo.Collection)

func StartDB() {

	client := GetClient()	
	// Check the connection
	err := client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func GetClient() *mongo.Client {
	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println(errEnv) 
	}
	mongoUrl := os.Getenv("MONGO_URL")
	fmt.Println(mongoUrl)
    clientOptions := options.Client().ApplyURI(mongoUrl)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    err = client.Connect(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    return client
}


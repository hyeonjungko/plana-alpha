package variables

import (
	"github.com/joho/godotenv"
	// "os"
	"log"
)


// var PORT string
// var MONGO_URL string
// var JWT_SECRET string


// var wg sync.WaitGroup

// var Secrets vars

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// PORT = os.Getenv("PORT")
	// MONGO_URL = os.Getenv("MONGO_URL")
	// JWT_SECRET = os.Getenv("JWT_SECRET")
}
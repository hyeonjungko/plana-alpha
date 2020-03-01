package main

//plana-kwxlv.run.goorm.io
import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"fmt"
	"os"


	"api/apiHandler"
	"api/mongoDB"
	// "api/variables"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	)

// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "home-tti")
// }


//NewRouter creates routers including various other type routers. 
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	apiHandler.AddAuthHandler(r)
	apiHandler.AddUserHandler(r)
	return r
}

func main() {
	router := NewRouter()
	mongoDB.StartDB()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	fmt.Println(PORT)

	log.Fatal(http.ListenAndServe(PORT, router))
}

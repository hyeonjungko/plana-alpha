package main

//plana-kwxlv.run.goorm.io
import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"os"


	"api/apiHandler"
	"api/mongoDB"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	// "github.com/dgrijalva/jwt-go"
)



// func home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "home-tti")
// }

// func signUp(w http.ResponseWriter, r *http.Request) {
// 	var newUser user
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Fprintf(w, "please enter user information")
// 	}

// 	json.Unmarshal(reqBody, &newUser)
// 	users = append(users, newUser)
// 	w.WriteHeader(http.StatusCreated)

// 	json.NewEncoder(w).Encode(newUser)
// }

// func logIn(w http.ResponseWriter, r *http.Request) {
// 	// pass
// 	return
// }



// func addMainHandler(r *mux.Router) {
// 	r.HandleFunc("/", home).Methods("GET")
// }


//NewRouter creates routers including various other type routers. 
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	apiHandler.AddAuthHandler(r)
	apiHandler.AddUserHandler(r)

	
	return r
}

func main() {
	mongoDB.StartDB()
	router := NewRouter()
	
	//getting PORT number from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	
	// router.HandleFunc("/", home).Methods("GET")
	

	log.Fatal(http.ListenAndServe(PORT, router))
}

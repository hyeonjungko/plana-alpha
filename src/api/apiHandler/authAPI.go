package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
	// "io/ioutil"
	"net/http"
	// "os"
	"api/auth"
	"api/mongoDB"
	"strconv"
	"time"
	// "reflect"

	// 	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

const (
	apiPathPrefix = "api/user"
)

type JWTTOKEN string

var dbAccess = mongoDB.NewUserAccess()

func AddAuthHandler(r *mux.Router) {
	r.HandleFunc("/login", apiLogInHandler).Methods("POST")
	r.HandleFunc("/signup", apiSignUpHandler).Methods("POST")
	r.HandleFunc("/logout", apiLogOutHandler).Methods("POST")
	r.HandleFunc("/check", apiCheckHandler).Methods("GET")
}

type Response struct {
	JWTToken string
	Error    error
}

//apiLogInHandler (1)gets login information (2)compare it with one in DB
func apiLogInHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := auth.PASSWORD(r.FormValue("password"))
	// fmt.Println([]byte(password))
	// fmt.Println(reflect.TypeOf(password))
	//calling login function to compare.
	user, err := dbAccess.LogIn(username, password)
	if err != nil {
		fmt.Fprint(w, "no user found")
	} else {
		fmt.Fprintf(w, fmt.Sprintf("Welcome %v\n", username))

		//generate token
		tokenString, tokenErr := user.GenerateToken(w, r)
		fmt.Fprintf(w, fmt.Sprintf("JWT Token: %s", tokenString))
		fmt.Println("JWT Token: ", tokenString)

		//send JWT token to response, so in frontend, we can access to the token.
		err = json.NewEncoder(w).Encode(Response{
			JWTToken: tokenString,
			Error:    tokenErr,
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
}

//apiSignUpHandler (1)gets new User (2)saves it to DB
func apiSignUpHandler(w http.ResponseWriter, r *http.Request) {
	//Unmarshalling, or convert json to golang data type
	id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	password := auth.PASSWORD(r.FormValue("password"))
	newUser := auth.User{
		Id:             auth.ID(id),
		Username:       username,
		FirstName:      firstname,
		LastName:       lastname,
		HashedPassword: password,
		SignedUpDate:   time.Now(),
	}
	//calling signup function to add newUser to database.
	fmt.Println(newUser)
	err2 := dbAccess.SignUp(newUser)
	fmt.Fprintf(w, fmt.Sprintf("%v\n", err2))
	fmt.Fprintf(w, fmt.Sprintf("%+v\n", newUser))
	//generate token
	tokenString, tokenErr := newUser.GenerateToken(w, r)
	fmt.Fprintf(w, fmt.Sprintf("JWT Token: %s", tokenString))
	fmt.Println("JWT Token: ", tokenString)

	//send JWT token to response, so in frontend, we can access to the token.
	err := json.NewEncoder(w).Encode(Response{
		JWTToken: tokenString,
		Error:    tokenErr,
	})
	if err != nil {
		log.Println(err)
		return
	}
}

func apiLogOutHandler(w http.ResponseWriter, r *http.Request) {

}

func apiCheckHandler(w http.ResponseWriter, r *http.Request) {

}

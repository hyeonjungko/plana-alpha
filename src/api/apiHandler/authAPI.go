package apiHandler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"api/auth"
	"api/mongoDB"
	"time"
	"os"
	// "reflect"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/joho/godotenv"
)

const (
	apiPathPrefix = "api/user"
)

var dbAccess = mongoDB.NewUserAccess()

func AddAuthHandler(r *mux.Router) {
	r.HandleFunc("/login", apiLogInHandler).Methods("POST")
	r.HandleFunc("/signup", apiSignUpHandler).Methods("POST")
	r.HandleFunc("/logout", apiLogOutHandler).Methods("POST")
	r.HandleFunc("/check", apiCheckHandler).Methods("GET")
}

//getting JWT_KEY from .env file
var errEnv = godotenv.Load()
var JwtKey = []byte(os.Getenv("JWT_SECRET"))

var expirationTime = time.Hour * 24

//claim to generate token
type UserClaim struct {
	Id       string
	Username string
	jwt.StandardClaims
}

//generateToken (1)method of user struct (2)returns JWT generated token (3)for authentication of usage in other private APIs.
func GenerateToken(w http.ResponseWriter, r *http.Request, username string, id string) (string, error) {
	expirationTime := time.Now().Add(expirationTime)
	claims := UserClaim{
		Id:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "민석이",
		},
	}
	// Sign and get the complete encoded token as a string using the secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Sorry, error while Signing Token!")
		return "", err
	}
	return tokenString, err
	// fmt.Printf("%v %v", tokenString, err)
}


type Response struct {
	JWTToken string
	Error    error
}

//apiLogInHandler (1)gets login information (2)compare it with one in DB
func apiLogInHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := auth.PASSWORD(r.FormValue("password"))
	userID, err := dbAccess.LogIn(username, password)
	if err != nil {
		fmt.Fprint(w, err)
	} else {
		fmt.Fprintf(w, fmt.Sprintf("Welcome %v\n", username))

		//generate token
		tokenString, tokenErr := GenerateToken(w, r, username, userID)
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
	// id, _ := strconv.Atoi(r.FormValue("id"))
	username := r.FormValue("username")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	password := auth.PASSWORD(r.FormValue("password"))
	//Joi 
	newUser := auth.User{
		Id:             primitive.NewObjectID(),
		Username:       username,
		FirstName:      firstname,
		LastName:       lastname,
		HashedPassword: password,
		SignedUpDate:   time.Now(),
	}
	
	//calling signup function to add newUser to database. gets userID from mongoDB
	userID, errSignup := dbAccess.SignUp(newUser)
	if errSignup != nil {
		fmt.Fprintf(w, fmt.Sprintf("%v\n", errSignup))
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("Welcome %v\n", username))
	//generate token
	tokenString, tokenErr := GenerateToken(w, r, username, userID)
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

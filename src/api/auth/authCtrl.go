//functions in this file controls the data in database.

package auth

import (
	"errors"
	// "golang.org/x/crypto/bcrypt"
	// "fmt"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"time"
)

//Casting Type
type ID int
type PASSWORD []byte

//Error
var ErrorMessage = errors.New("No user found")

type User struct {
	Id             ID        `json:"user_id"`
	Username       string    `json:"username"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName`
	HashedPassword PASSWORD  `json:"hashedPassword"`
	SignedUpDate   time.Time `json:"signedUpDate,omitempty"`
}

//slice of user structs
var users = []User{
	{
		Id:             1,
		Username:       "hko",
		FirstName:      "Hyeonjung",
		LastName:       "Ko",
		HashedPassword: PASSWORD("asdfjsdkl123a"),
		SignedUpDate:   time.Now(),
	},
}

//userAccess is an Interface for user database requires
type UserAccess interface {
	SignUp(u User) error                                    //signup
	LogIn(userName string, password PASSWORD) (User, error) //login
	// LogOut() error                          //logout
	// Check(Id ID) (user, error)              //check if user is loggedin
}

//getting JWT_KEY from .env file
var err = godotenv.Load()
var JwtKey = []byte(os.Getenv("JWT_SECRET"))

var expirationTime = time.Hour * 24

//claim to generate token
type UserClaim struct {
	Id       ID
	Username string
	jwt.StandardClaims
}

//generateToken (1)method of user struct (2)returns JWT generated token (3)for authentication of usage in other private APIs.
func (u *User) GenerateToken(w http.ResponseWriter, r *http.Request) (string, error) {
	expirationTime := time.Now().Add(expirationTime)
	claims := UserClaim{
		Id:       u.Id,
		Username: u.Username,
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



package main 
import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"time"
)


//Casting Type
type ID int
type PASSWORD []byte

//Error
var ErrorMessage = errors.New("No user found")


type user struct {
	Id             ID `json:"user_id"`
	Username       string `json:"username"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName`
	HashedPassword PASSWORD `json:"hashedPassword"`
	SignedUpDate  time.Time `json:"signedUpDate"`
}


//slice of user structs
var users = []user{
	{
		Id:             1,
		Username:       "hko",
		FirstName:      "Hyeonjung",
		LastName:       "Ko",
		HashedPassword: PASSWORD("asdfjsdkl123a"),
	},
}

//userAccess is an Interface for user database requires
type userAccess interface {
	SignUp(u user) error                //signup
	LogIn(userName string, password PASSWORD) (user, error)  //login
	LogOut() error                          //logout
	Check(Id ID) (user, error)              //check if user is loggedin
}

//instant database 
type instantDBAccess struct {
	users []user
}

//
func NewInstantDBAccess() userAccess {
	return &instantDBAccess{
		users: users,
	}
}

//signUp (1)gets user struct (2)convert password to hashcode (3)include into user database
func (dbAccess *instantDBAccess) SignUp(u user) error {
	hashedPassword, err := bcrypt.GenerateFromPassword(u.HashedPassword, bcrypt.DefaultCost)
	// fmt.Println(hashedPassword, err)
	if err != nil {
		return ErrorMessage
	}
	u.HashedPassword = hashedPassword
	dbAccess.users = append(dbAccess.users, u)
	return fmt.Errorf("SignUp Complete! Welcome %s!", u.Username)
}

//login gets userName and password to compare with data in DB
func (dbAccess *instantDBAccess) LogIn(userName string, password PASSWORD) (user, error) {
	for _, user := range dbAccess.users {
		if user.Username == userName {
			// fmt.Println(user.HashedPassword)
			err := bcrypt.CompareHashAndPassword(user.HashedPassword, password) 
			if err == nil {
				return user, fmt.Errorf("Login Success! user: %s", user.Username)
			}
		} 
	}
	return user{}, ErrorMessage
}

//logout removes the current user's JWT token of browser
func (dbAccess *instantDBAccess) LogOut() error {
	return errors.New("Logged out successfully")
}

//check returns current user struct
func (dbAccess *instantDBAccess) Check(Id ID) (user, error) {
	for _, user := range dbAccess.users {
		if(user.Id == Id){
			return user, fmt.Errorf("Current User: %s", user.Username)
		}
	}
	return user{}, ErrorMessage
}

func main() {
	dbAccess := NewInstantDBAccess()
	dbAccess.SignUp(user{
		Id:        ID(2),
		Username:  "felipe",
		FirstName: "Minseok",
		LastName:  "Choi",
		HashedPassword: PASSWORD("1q2w3e4"),
		SignedUpDate: time.Now(),
	})
	// fmt.Println(reflect.TypeOf(time.Now()))
	// fmt.Println(dbAccess)
}

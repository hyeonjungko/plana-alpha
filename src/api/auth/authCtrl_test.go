package auth

import (
	"fmt"
	"time"
	   )

func ExampleNewInstantDBAccess(){
	//Creating instance of DB strct
	dbAccess := NewInstantDBAccess()
	
	//SignUp returns error
	err1 := dbAccess.SignUp(user{
		Id:        ID(2),
		Username:  "felipe",
		FirstName: "Minseok",
		LastName:  "Choi",
		HashedPassword: PASSWORD("1q2w3e4r"),
		SignedUpDate: time.Now(),
	})
	fmt.Println(err1)
	
	//LogIn returns loggedin user and error
	_, err2 := dbAccess.LogIn("felipe", PASSWORD("1q2w3e4r"))
	fmt.Println(err2)
	
	//Check returns current user and error
	_, err3 := dbAccess.Check(ID(2))
	fmt.Println(err3)
	
	//LogOut returns error
	err4 := dbAccess.LogOut()
	fmt.Println(err4)
	
	//Output:
	//SignUp Complete! Welcome felipe!
	//Login Success! user: felipe
	//Current User: felipe
	//Logged out successfully
}

//instant database
// type instantDBAccess struct {
// 	Users []User
// }

// func NewInstantDBAccess() userAccess {
// 	return &instantDBAccess{
// 		Users: users,
// 	}
// }

//signUp (1)gets user struct (2)convert password to hashcode (3)include into user database
/*func (dbAccess *instantDBAccess) SignUp(u user)  error {
	hashedPassword, err := bcrypt.GenerateFromPassword(u.HashedPassword, bcrypt.DefaultCost)
	if err != nil {
		return ErrorMessage
	}
	u.HashedPassword = hashedPassword
	dbAccess.Users = append(dbAccess.Users, u)
	return fmt.Errorf("SignUp Complete! Welcome %s!", u.Username)
}

//login gets userName and password to compare with data in DB
func (dbAccess *instantDBAccess) LogIn(userName string, password PASSWORD) (user, error) {

	for _, user := range dbAccess.Users {
		if user.Username == userName {
			err := bcrypt.CompareHashAndPassword(user.HashedPassword, password)
			if err == nil {
				return user, nil //fmt.Errorf("Login Success! user: %s", user.Username)
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
	for _, user := range dbAccess.Users {
		if(user.Id == Id){
			return user, fmt.Errorf("Current User: %s", user.Username)
		}
	}
	return user{}, ErrorMessage
}
*/



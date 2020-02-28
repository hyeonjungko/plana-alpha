package main

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




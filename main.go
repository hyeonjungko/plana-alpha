package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type location struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type user struct {
	ID        string `json:"user_id"`
	Username  string `json:"username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName`
}

type allLocations []location
type allUsers []user

var locations = allLocations{
	{
		ID:          "1",
		Name:        "Aquamarine",
		Description: "Cafe by the beach",
	},
}

var users = allUsers{
	{
		ID:        "1",
		Username:  "hko",
		FirstName: "Hyeonjung",
		LastName:  "Ko",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home-tti")
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "please enter user information")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

func logIn(w http.ResponseWriter, r *http.Request) {
	// pass
	return
}

func userPlans(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	// fetch user plans from DB
	return
}

func userLikes(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	// fetch user likes from DB
	return
}

func userSettings(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	// fetch user settings from DB
	return
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userId"]

	// fetch user profile from DB
	return
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var newLocation location
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter location name and description")
	}
	json.Unmarshal(reqBody, &newLocation)
	locations = append(locations, newLocation)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newLocation)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/user/login", logIn).Methods("POST")
	router.HandleFunc("/user/signup", signUp).Methods("POST")
	router.HandleFunc("/user/{userId}/plans", userPlans).Methods("GET")
	router.HandleFunc("/user/{userId}/likes", userLikes).Methods("GET")
	router.HandleFunc("/user/{userId}/settings", userSettings).Methods("GET")
	router.HandleFunc("/user/{userId}/profile", userProfile).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

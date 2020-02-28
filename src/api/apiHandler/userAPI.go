package apiHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"api/auth"

	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func AddUserHandler(router *mux.Router) {
	userRouter := mux.NewRouter().StrictSlash(true)
	userRouter.HandleFunc("/user/{userId}/plans", userPlans).Methods("GET")
	userRouter.HandleFunc("/user/{userId}/likes", userLikes).Methods("GET")
	userRouter.HandleFunc("/user/{userId}/settings", userSettings).Methods("GET")
	userRouter.HandleFunc("/user/{userId}/profile", userProfile).Methods("GET")
	router.PathPrefix("/user").Handler(negroni.New(
		negroni.HandlerFunc(JwtMiddleware.HandlerWithNext),
		negroni.Wrap(userRouter),
	))
}

type allLocations []location

var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

type location struct {
	ID          string `json:"iD"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var locations = allLocations{
	{
		ID:          "1",
		Name:        "Aquamarine",
		Description: "Cafe by the beach",
	},
}

func userPlans(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	userInfo := user.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(userInfo)
	// userID := mux.Vars(r)["userId"]
	fmt.Fprintf(w, fmt.Sprintf("This is plans! with %v", userInfo))
	// fetch user plans from DB
	return
}

func userLikes(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	userInfo := user.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(userInfo)
	// userID := mux.Vars(r)["userId"]
	fmt.Fprintf(w, fmt.Sprintf("This is Likes! with %v", userInfo))
	// userID := mux.Vars(r)["userId"]

	// fetch user likes from DB
	return
}

func userSettings(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	userInfo := user.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(userInfo)
	// userID := mux.Vars(r)["userId"]
	fmt.Fprintf(w, fmt.Sprintf("This is Settings! with %v", userInfo))
	// userID := mux.Vars(r)["userId"]

	// fetch user settings from DB
	return
}

func userProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	userInfo := user.(*jwt.Token).Claims.(jwt.MapClaims)
	fmt.Println(userInfo)
	// userID := mux.Vars(r)["userId"]
	fmt.Fprintf(w, fmt.Sprintf("This is Profile! with %v", userInfo))
	// userID := mux.Vars(r)["userId"]

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

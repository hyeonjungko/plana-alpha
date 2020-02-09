package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Location struct {
	ID          string `json:"ID"`
	Name        string `json:"Name"`
	Description string `json:"Description"`
}

type allLocations []Location

var locations = allLocations{
	{
		ID:          "1",
		Name:        "Aquamarine",
		Description: "Cafe by the beach",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is home!")
}

func createLocation(w http.ResponseWriter, r *http.Request) {
	var newLocation Location
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
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8080", router))
}

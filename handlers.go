package main

import (
	"fmt"
	"net/http"
	"time"

	//"encoding/json"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func FileSHA256(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sha256 := vars["sha256"]

	fmt.Fprintln(w, "SHA256:", sha256)

	// TODO: Lookup cache before database

	var result Reputation
	result = lookupReputation(sha256)

	fmt.Fprintln(w, "Reputation:", result.Rep)

	// TODO: Update cache

}

func InsertFileRep(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sha256 := vars["sha256"]
	rep := vars["rep"]
	ttl := 10

	fmt.Fprintln(w, "File Hash:", sha256)
	fmt.Fprintln(w, "Reputation:", rep)

	// For now we just return a canned response
	/*newRep := Reputations{
		Reputation{Value: "good", DateAdded: "", TTL: ttl},
	}*/

	currentTime := time.Now()
	newRep := Reputation{SHA256: sha256, Rep: rep, DateAdded: currentTime.String(), TTL: ttl}

	insertReputation(newRep)

}

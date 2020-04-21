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

	// Check the cache to see if there is an entry for this key
	checkCacheTime := time.Now()
	item, found := globalcache.Get(sha256)
	if found {
		fmt.Fprintln(w, "Cache hit, reputation is: ", item)
		fmt.Fprintln(w, "Cache lookup time: ", time.Since(checkCacheTime).Microseconds())
		return
	}
	fmt.Fprintln(w, "No cache hit, need to lookup MongoDB.")

	var result Reputation
	result = lookupReputation(sha256)

	fmt.Fprintln(w, "Reputation: ", result.Rep)

	// Update cache with this entry
	// TODO: Update existing entry if it already exists (or validate that zizou works like this)
	globalcache.Set(sha256, result.Rep, 10*time.Second) // TODO: Fix TTL value

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

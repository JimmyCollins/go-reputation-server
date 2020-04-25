package main

import (
	"fmt"
	"net/http"
	"time"

	//"encoding/json"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome! Nothing to see here.")
}

func FileSHA256(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sha256 := vars["sha256"]

	fmt.Fprintln(w, "SHA256:", sha256)

	// Check the cache to see if there is an entry for this key
	cacheLatency := time.Now()
	item, found := globalcache.Get(sha256)
	if found {
		fmt.Fprintln(w, "Cache hit, reputation is: ", item)
		fmt.Fprintln(w, "Cache lookup time (ms): ", time.Since(cacheLatency).Milliseconds())
		return
	}
	fmt.Fprintln(w, "No cache hit, need to lookup MongoDB.")

	var result Reputation
	dbLatency := time.Now()
	result = lookupReputation(sha256)
	fmt.Fprintln(w, "Database lookup time (ms): ", time.Since(dbLatency).Milliseconds())

	fmt.Fprintln(w, "Reputation: ", result.Rep)

	// Update cache with this entry
	globalcache.Set(sha256, result.Rep, time.Duration(result.TTL)*time.Second)

}

func InsertFileRep(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sha256 := vars["sha256"]
	rep := vars["rep"]
	ttl := 10 // Default timeout is 10s - TODO: Externalize

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

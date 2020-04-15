package main

import (
	"fmt"
	"net/http"
	"time"

	//"encoding/json"
	vt "github.com/VirusTotal/vt-go"
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
		fmt.Fprintln(w, "Cache Hit: MD5 of file: ", item)
		fmt.Fprintln(w, "Cache Time: ", time.Since(checkCacheTime))
		return
	}
	fmt.Fprintln(w, "No Cache Hit: Cache Time: ", time.Since(checkCacheTime))

	// Note: You need a valid VirusTotal API key for this to work
	// You can sign-up for a limited access public account for free
	var apikey = ""

	// Lookup the VirusTotal API to the MD5 of this file, using the SHA256 as the key
	client := vt.NewClient(apikey)
	apiTime := time.Now()
	file, err := client.GetObject(vt.URL("files/%s", sha256))
	if err != nil {
		fmt.Fprintln(w, "Error: ", err)
	}
	fmt.Fprintln(w, "API Request Time: ", time.Since(apiTime))

	ls, err := file.GetTime("last_submission_date")
	if err != nil {
		fmt.Fprintln(w, "Error: ", err)
	}

	md5, err := file.GetString("md5")
	if err != nil {
		fmt.Fprintln(w, "Error: ", err)
	}

	fmt.Fprintln(w, "MD5 is: ", md5)
	fmt.Fprintln(w, "\nFile was submitted for the last time on\n", ls)

	// Update the cache, add a TTL of 10 seconds (for testing)
	globalcache.Set(sha256, md5, 10*time.Second)

	// For now we just return a canned response
	/*rep := Reputations{
	    Reputation{Value: "good", TTL: 15},
	}*/

	//json.NewEncoder(w).Encode(rep)

	/*if err := json.NewEncoder(w).Encode(rep); err != nil {
	    panic(err)
	}*/

}

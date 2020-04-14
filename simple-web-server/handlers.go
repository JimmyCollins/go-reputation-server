package main
 
import (
    "fmt"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    vt "github.com/VirusTotal/vt-go"
)
 

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}
 

func FileSHA256(w http.ResponseWriter, r *http.Request) {

    vars := mux.Vars(r)
    sha256 := vars["sha256"]

    fmt.Fprintln(w, "SHA256:", sha256)

    // Note: You need a valid VirusTotal API key for this to work
    // You can sign-up for a limited access public account for free
    var apikey = "455f18250d4b09c1bcf67e5797a2831cae76ae24e3fe37432381ce7dff539686";

    client := vt.NewClient(apikey)

    file, err := client.GetObject(vt.URL("files/%s", sha256))
	if err != nil {
		fmt.Fprintln(w, "Error: ", err);
    }

    ls, err := file.GetTime("last_submission_date")
	if err != nil {
		fmt.Fprintln(w, "Error: ", err);
	}

	fmt.Fprintln(w, "\nFile was submitted for the last time on\n", file.ID(), ls)
    
    // For now we just return a canned response
    /*rep := Reputations{
        Reputation{Value: "good", TTL: 15},
    }*/

    //json.NewEncoder(w).Encode(rep)

    /*if err := json.NewEncoder(w).Encode(rep); err != nil {
        panic(err)
    }*/

}
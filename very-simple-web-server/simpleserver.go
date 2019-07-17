package main
 
import (
    "fmt"
    "log"
    "net/http"
 	"encoding/json"

 	// Used for routing
    "github.com/gorilla/mux"
)
 

type Reputation struct {
    Value      string	`json:"name"`
    TTL int             `json:"ttl"`
}
 
type Reputations []Reputation



func main() {
 
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    router.HandleFunc("/file/{md5}", FileMD5)
 
    log.Fatal(http.ListenAndServe(":8080", router))
}
 


func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome! There's nothing here to see.")
}
 


func FileMD5(w http.ResponseWriter, r *http.Request) {
    //vars := mux.Vars(r)
    //md5 := vars["md5"]
    //fmt.Fprintln(w, "MD5:", md5)

    //fdbf231e5f33f31199bd95d3fee35e50

    rep := Reputations{
        Reputation{Value: "good", TTL: 15},
    }

    json.NewEncoder(w).Encode(rep)

}
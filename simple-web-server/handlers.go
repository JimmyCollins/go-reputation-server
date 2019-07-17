package main
 
import (
    "fmt"
    "net/http"
 	"encoding/json"
)
 

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

    //json.NewEncoder(w).Encode(rep)

    if err := json.NewEncoder(w).Encode(rep); err != nil {
        panic(err)
    }

}
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	zz "github.com/arriqaaq/zizou"
)

var globalcache *zz.Cache

func main() {

	router := NewRouter()

	// Setup the in-memory cache
	initializeCache()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeCache() {
	config := zz.Config{
		SweepTime: 10 * time.Minute,
		ShardSize: 256,
	}

	var err error

	globalcache, err = zz.New(&config)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Cache Setup done")
}

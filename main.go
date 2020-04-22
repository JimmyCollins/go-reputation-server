package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	zz "github.com/arriqaaq/zizou"
)

var globalcache *zz.Cache

func main() {

	router := NewRouter()

	// Setup the in-memory cache
	initializeCache()

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)

	// Server setup
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
	}

	// Start server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
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

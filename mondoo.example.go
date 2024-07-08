package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mondoo Engineer!")
	})

	srv := &http.Server{
		Addr:         ":80",
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

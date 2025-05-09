package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Starting serve on :8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

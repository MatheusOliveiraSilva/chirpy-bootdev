package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("./app"))))
	mux.Handle("/app/assets/", http.StripPrefix("/app/assets/", http.FileServer(http.Dir("./assets"))))

	mux.HandleFunc("/healthz", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("OK"))
	})

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

package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go program!"))
}

func main() {
	server := http.NewServeMux()
	
	server.HandleFunc("/hello", handleRoot)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8000", server)
	if err != nil {
		fmt.Println("Error while opening the server:", err)
		os.Exit(1)
	}
}
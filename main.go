package main

import (
	"fmt"
	"net/http"
	"os"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hellow from Go program!"))
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/", handleRoot)

	err := http.ListenAndServe(":8000", server)
	if err == nil{
		fmt.Println("Error while opening the server")
		os.Exit(1)
	}
}
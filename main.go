package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", 
	func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hellow from Go program"))
	})
	
	err := http.ListenAndServe(":8000", nil)
	if err == nil{
		fmt.Println("Error while opening the server")
		os.Exit(1)
	}
}
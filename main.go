package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside Handler")
	fmt.Fprint(w, "Welcome to the  Obingo77 Api!!!"+req.URL.Path[1:])
}

func aboutMe(w http.ResponseWriter, req *http.Request) {
	who := "Obrien Makenzi"

	fmt.Fprint(w, "App made by OM <3...")
	fmt.Println("Endpoint hit:", who)
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/aboutMe", aboutMe)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err.Error)
	}
}

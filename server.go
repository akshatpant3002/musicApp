package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port string = ":8000"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Router is running...")
	})
	router.HandleFunc("/groupPost", getGroups).Methods("GET") // might need to change with option
	router.HandleFunc("/groupPost", addGroups).Methods("POST", "OPTIONS")
	router.HandleFunc("/userPost", getUsers).Methods("GET") // might need to change with option
	router.HandleFunc("/userPost", addUsers).Methods("POST", "OPTIONS")

	log.Print("Server listening on port: ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}

//router.HandleFunc("/groupPost", addGroups).Methods("POST")

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthCheckHandler)

	r.HandleFunc("/students", getAllStudentsHandler)
	r.HandleFunc("/student/{id}", getStudentByIDHandler)

	fmt.Println("Server listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}
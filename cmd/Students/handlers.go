package main

import (
	"encoding/json"
	"net/http"

	students "github.com/Skapar/go-tsis1/pkg/Students"
	"github.com/gorilla/mux"
)

type Response struct {
	Students []students.Student `json:"student"`
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": students.Info()})
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	restaurants := students.GetRestaurants()
	respondWithJSON(w, http.StatusOK, restaurants)
}

func restaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	restaurant, err := students.GetRestaurant(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, restaurant)
}
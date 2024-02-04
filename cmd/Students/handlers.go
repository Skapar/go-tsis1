package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	students "github.com/Skapar/go-tsis1"
)

type Response struct {
	Restaurants []students.Restaurant `json:"restaurants"`
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
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": abrplus.Info()})
}

func restaurants(w http.ResponseWriter, r *http.Request) {
	restaurants := abrplus.GetRestaurants()
	respondWithJSON(w, http.StatusOK, restaurants)
}

func restaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	restaurant, err := abrplus.GetRestaurant(id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, restaurant)
}
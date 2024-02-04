package main

import (
	"encoding/json"
	"net/http"

	studentik "github.com/Skapar/go-tsis1/pkg/students"
	"github.com/gorilla/mux"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type HealthCheckResponse struct {
	Status string `json:"status"`
	Check   string `json:"Check"`
}

func writeJSONResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(ErrorResponse{"Internal Server Error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	writeJSONResponse(w, http.StatusOK, HealthCheckResponse{"ok", studentik.Check()})
}

func getAllStudentsHandler(w http.ResponseWriter, r *http.Request) {
	allStudents := studentik.GetAllStudents()
	writeJSONResponse(w, http.StatusOK, allStudents)
}

func getStudentByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentID := vars["id"]

	student, err := studentik.GetStudentByID(studentID)
	if err != nil {
		writeJSONResponse(w, http.StatusNotFound, ErrorResponse{"Student not found"})
		return
	}

	writeJSONResponse(w, http.StatusOK, student)
}
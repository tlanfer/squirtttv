package config

import (
	"companion/internal/config"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			get(w, r)
		case http.MethodPost:
			post(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func get(w http.ResponseWriter, _ *http.Request) {
	dto := config.Get()
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(dto)
	if err != nil {
		log.Println("Error encoding json", err)
		http.Error(w, "Error encoding json", http.StatusInternalServerError)
		return
	}
}

func post(w http.ResponseWriter, req *http.Request) {
	var dto config.Config
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		log.Println("Error decoding json", err)
		http.Error(w, "Error decoding json", http.StatusBadRequest)
		return
	}
	config.Set(dto)
	w.WriteHeader(http.StatusAccepted)
}

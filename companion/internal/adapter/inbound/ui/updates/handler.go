package updates

import (
	"companion/internal/adapter/outbound/update"
	"encoding/json"
	"log"
	"net/http"
)

func NewHandler(version string) http.Handler {
	return &handler{
		version: version,
	}
}

type handler struct {
	version string
}

func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.get(w, req)
	case http.MethodPost:
		h.post(w, req)
	}
}

func (h handler) get(w http.ResponseWriter, _ *http.Request) {
	latest, b, err := update.IsLatest()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dto := dto{
		Version:  h.version,
		Latest:   latest,
		IsLatest: b,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println("error encoding response", err)
	}
}

func (h handler) post(w http.ResponseWriter, _ *http.Request) {
	log.Println("would update now")
	err := update.Run()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

type dto struct {
	Version  string `json:"version"`
	Latest   string `json:"latest"`
	IsLatest bool   `json:"isLatest"`
}

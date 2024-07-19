package updates

import (
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
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h handler) get(w http.ResponseWriter, _ *http.Request) {

	dto := dto{
		Version: h.version,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		log.Println("error encoding response", err)
	}
}

type dto struct {
	Version string `json:"version"`
}

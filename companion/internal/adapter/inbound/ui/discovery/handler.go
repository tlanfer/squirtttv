package discovery

import (
	"companion/adapter/outbound/scanner"
	"encoding/json"
	"net/http"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ = json.NewEncoder(w).Encode(dto{
			Hosts: scanner.List(),
		})
	})
}

type dto struct {
	Hosts []string `json:"hosts"`
}

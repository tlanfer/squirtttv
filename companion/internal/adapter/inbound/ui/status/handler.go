package status

import (
	"companion/internal/state"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

func NewHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			get(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func get(w http.ResponseWriter, r *http.Request) {

	headerWait := parseBool(r.FormValue("wait"))
	headerTwitch := parseBool(r.FormValue("twitch"))
	headerStreamlabs := parseBool(r.FormValue("streamlabs"))
	headerStreamelements := parseBool(r.FormValue("streamelements"))

	timeout := time.Now().Add(10 * time.Second)
	if headerWait {
		for {
			if state.TwitchConnected != headerTwitch {
				break
			}
			if state.StreamlabsConnected != headerStreamlabs {
				break
			}
			if state.StreamElementsConnected != headerStreamelements {
				break
			}

			if time.Now().After(timeout) {
				break
			}

			time.Sleep(500 * time.Millisecond)
		}
	}

	d := dto{
		Twitch:         state.TwitchConnected,
		Streamlabs:     state.StreamlabsConnected,
		Streamelements: state.StreamElementsConnected,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(d)
	if err != nil {
		log.Println("Error encoding json", err)
	}
}

type dto struct {
	Twitch         bool `json:"twitch"`
	Streamlabs     bool `json:"streamlabs"`
	Streamelements bool `json:"streamelements"`
}

func parseBool(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}

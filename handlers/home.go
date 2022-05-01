package handlers

import (
	"encoding/json"
	"github.com/PlatziGoRestWebsockets/server"
	"net/http"
)

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode(HomeResponse{
			Message: "Welcome to Home",
			Status:  true,
		})
		if err != nil {
			return
		}
	}
}
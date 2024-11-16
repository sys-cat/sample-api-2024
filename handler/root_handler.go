package handler

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type RootResponse struct {
	Message string `json:"message"`
}

func RootIndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := json.NewEncoder(w).Encode(RootResponse{Message: "Hello, World!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

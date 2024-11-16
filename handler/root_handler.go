package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type RootResponse struct {
	Message string `json:"message"`
}

func RootIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	message := "Hello, World!"
	if len(p) == 0 {
		logger.Info("no params", "params", p)
	}
	if p.ByName("message") != "" {
		logger.Info("one param", "params", p)
		message = p.ByName("message")
	}
	err := json.NewEncoder(w).Encode(RootResponse{Message: message})
	if err != nil {
		logger.Error("json encode error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

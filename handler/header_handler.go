package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type HeaderResponse struct {
	Method  string
	Path    string
	Headers []Header
}

type Header struct {
	Index string
	Value []string
}

func HeaderIndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Param) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("request", "uri", r.RequestURI)
	headers := []Header{}
	for key, values := range r.Header {
		headers = append(headers, Header{
			Index: key,
			Value: values,
		})
	}
	res := HeaderResponse{
		Method:  r.Method,
		Path:    r.URL.Path,
		Headers: headers,
	}
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		logger.Error("json encode error", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

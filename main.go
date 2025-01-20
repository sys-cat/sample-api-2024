package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/sys-cat/sample-api-2024/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := httprouter.New()
	router.GET("/api/v1/sample", handler.InfoHandler)
	router.GET("/", handler.RootIndexHandler)
	router.GET("/sample-api/", handler.RootIndexHandler)
	router.GET("/sample-api/:message", handler.RootIndexHandler)
	router.GET("/health", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(http.StatusOK)
	})
	logger.Info("server start", "port", ":8080")
	logger.Error("server start", "error", http.ListenAndServe(":8080", router))
}

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
	router.GET("/", handler.RootIndexHandler)
	router.GET("/:message", handler.RootIndexHandler)

	logger.Error("server start", "error", http.ListenAndServe(":8080", router))
}

package main

import (
	"net/http"
	"local/sample-api-20024/handler"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", handler.RootIndexHandler)

	http.ListenAndServe(":8080", router)
}

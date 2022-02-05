package main

import (
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/logger"
	"net/http"
)

func main() {

	bootstrap.SetupDB()
	// gorilla/mux 因实现了 net/http 包的 http.Handler 接口，故兼容 http.ServeMux
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

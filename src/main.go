package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server"
)

func init() {
	router := gin.New()
	server.ApiEntry(router)
	http.Handle("/", router)
}

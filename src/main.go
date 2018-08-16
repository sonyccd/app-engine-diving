package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server"
)

func init() {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	server.ApiEntry(router)
	http.Handle("/", router)
}

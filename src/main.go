package main

import (
	"google.golang.org/appengine"
	"net/http"

	"github.com/gin-gonic/gin"
	"server"
)

func main() {
	router := gin.New()
	router.Use(gin.ErrorLogger())
	router.Use(gin.Recovery())
	server.ApiEntry(router)
	http.Handle("/", router)

	appengine.Main()
}

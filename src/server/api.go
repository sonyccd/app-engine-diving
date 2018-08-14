package server

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func ApiEntry(router *gin.Engine) {
	appApi := router.Group("/api")
	appApi.Use(ApiInjector())

	testApi := appApi.Group("/test")
	testApi.GET("/", api.Ping)

	diveGroup := appApi.Group("/dive")
	diveGroup.GET("/")
	diveGroup.GET("/:id", api.DiveGet)
	diveGroup.POST("/", api.DiveCreate)

	userGroup := appApi.Group("/user")
	userGroup.GET("/:uid", api.CreateUpdateUser)

	authGroup := appApi.Group("/auth")
	authGroup.GET("/", api.GetAuth)
}

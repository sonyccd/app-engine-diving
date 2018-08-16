package server

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

func ApiEntry(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.File("/Users/bradfordbazemore/Devel/hermes/hermes-app-engine/src/client/index.html")
	})

	appApi := router.Group("/api")
	appApi.Use(SetupSession(), ApiInjector())
	{
		testApi := appApi.Group("/test")
		testApi.GET("/", api.Ping)

		diveGroup := appApi.Group("/dive")
		diveGroup.GET("/", api.TestDive)
		diveGroup.GET("/:id", api.DiveGet)
		diveGroup.POST("/", api.DiveCreate)

		userGroup := appApi.Group("/user")
		userGroup.GET("/:uid", api.GetUser)

		authGroup := appApi.Group("/auth")
		authGroup.GET("/", api.GetAuth)
	}
}

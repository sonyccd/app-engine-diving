package server

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"model"
)

func Injector() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := appengine.NewContext(c.Request)
		c.Set("DiveInterface", model.NewDiveImplementation(ctx))
	}
}

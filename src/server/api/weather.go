package api

import (
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
)

func WeatherGet(c *gin.Context) {
	if weatherInt, found := c.Keys["WeatherInterface"].(model.WeatherInterface); !found {
		c.String(http.StatusNotFound, "")
		return
	} else {
		weather := weatherInt.Pull()
		c.String(http.StatusOK, weather.Summary)
	}
}

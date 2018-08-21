package model

import (
	"context"
	"firebase.google.com/go"
	"github.com/shawntoffel/darksky"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

type Weather struct {
}

type WeatherInterface interface {
	Pull() *darksky.DataPoint
}

type WeatherImplementation struct {
	appCtx      context.Context
	firebaseApp *firebase.App
}

func NewWeatherImplementation(appEngineCtx context.Context, firebaseApp *firebase.App) WeatherInterface {
	return WeatherImplementation{appCtx: appEngineCtx, firebaseApp: firebaseApp}
}

func (w WeatherImplementation) Pull() *darksky.DataPoint {
	client := darksky.NewWithClient("0572265ceb918ccb26859277f09bde18", urlfetch.Client(w.appCtx))

	request := darksky.ForecastRequest{}
	request.Latitude = 40.7128
	request.Longitude = -74.0059
	request.Options = darksky.ForecastRequestOptions{Exclude: "hourly,minutely"}

	response, err := client.Forecast(request)

	if err != nil {
		log.Errorf(w.appCtx, err.Error())
		return nil
	}

	return response.Currently
}

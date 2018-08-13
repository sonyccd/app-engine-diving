package api

import (
	"firebase.google.com/go"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
	"net/http"
)

func GetAuth(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile("/Users/bradfordbazemore/Devel/hermes/hermes-app-engine/project-hermes-staging-firebase-adminsdk-q2yxf-fd6ecd39e6.json"))
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	if app == nil {
		c.String(http.StatusNotFound, "Could not init firebase app")
		return
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	defer client.Close()
	iter := client.Collection("Dive").Documents(ctx)
	temp := "data:"
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		temp = fmt.Sprint(temp, doc.Data())
	}
	c.String(http.StatusOK, temp)
}

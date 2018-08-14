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

var token = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjY3NmY5MmU1MGQ5ZmUxNzdiM2I5NTJjNGM4ZWU2YjY1ZDk3ZWIwZmMifQ.eyJpc3MiOiJodHRwczovL3NlY3VyZXRva2VuLmdvb2dsZS5jb20vcHJvamVjdC1oZXJtZXMtc3RhZ2luZyIsIm5hbWUiOiJCcmFkIEJhemVtb3JlIiwicGljdHVyZSI6Imh0dHBzOi8vbGg0Lmdvb2dsZXVzZXJjb250ZW50LmNvbS8tU0gwckRCelV2NkUvQUFBQUFBQUFBQUkvQUFBQUFBQUFFOU0vN2l5ek5Tb1BETE0vcGhvdG8uanBnIiwiYXVkIjoicHJvamVjdC1oZXJtZXMtc3RhZ2luZyIsImF1dGhfdGltZSI6MTUzNDIwOTcwNSwidXNlcl9pZCI6IkZZQ1JCbzlpVFZRNHBnYnNTYWpJZWhraHhYWDIiLCJzdWIiOiJGWUNSQm85aVRWUTRwZ2JzU2FqSWVoa2h4WFgyIiwiaWF0IjoxNTM0MjA5NzA1LCJleHAiOjE1MzQyMTMzMDUsImVtYWlsIjoiYmFua2Vyc2FsZ29AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMDkxNjQ5MjI5MjcyNzI0MTcxMTQiXSwiZW1haWwiOlsiYmFua2Vyc2FsZ29AZ21haWwuY29tIl19LCJzaWduX2luX3Byb3ZpZGVyIjoiZ29vZ2xlLmNvbSJ9fQ.oNbXAbb8uf3FC4GMYwPojxaqk3EJs09FnjiLSzX5pxwZTeRzJdKDw0lQelUF4gvuSeMgMxerVZwnALBt_H7QXxHoHLVibsaBL-n50ic3VspFmOE6XMmJu02rv1gIh9rLQ0jCAhwhQYqTSkcUBk5Nb4MuLGu6k1KBOdBdK7KSNUs2ji24xfKfEbu37zOjARFIzhO-ac2IFPZuVuOiv_CDYFsqwZ2wPVi5t2ZV--o7Bc-cFPnZ02_CKHSxaEwFaT7pfaPJXj90uIJ51rolcYeKp0XMKINLglq2BVH0CVwLrLiVu_IoJ-T3QctTrMViV-iJafSSft8tJmjEJaLkDp9S_w"

func FirestoreTest(c *gin.Context) {
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
	client, err := app.Auth(ctx)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	t, err := client.VerifyIDToken(ctx, token)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	user, err := client.GetUser(ctx, t.UID)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user.UserInfo)
}

package http

import (
	"context"
	"github.com/gin-gonic/gin"
	go_gin "golang.source-fellows.com/seminar/go-gin/v2"
	"log"
	"net/http"
)

func StartServer(repository go_gin.AutoRepository) error {
	r := gin.Default()
	r.GET("/ping", handlePing)

	authorized := r.Group("/api", gin.BasicAuth(gin.Accounts{
		"foo": "bar",
	}))
	authorized.GET("/getautos", handleGetAutos(repository))
	authorized.PUT("/addauto", handleAddAuto(repository))
	return r.Run()
}

func handleAddAuto(repository go_gin.AutoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		audi := &go_gin.Audi{}
		err := c.BindJSON(audi)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		repository.AddAuto(c.Request.Context(), audi)
	}
}

func handlePing(context *gin.Context) {
	context.Data(http.StatusOK, "text/plain", []byte("pong"))
}

func handleGetAutos(repository go_gin.AutoRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceid := c.GetHeader("x-trace-id")
		ctx := context.WithValue(c.Request.Context(), "x-trace-id", traceid)
		autos, err := repository.GetAllAuto(ctx)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		c.Negotiate(http.StatusOK, gin.Negotiate{
			Offered: []string{gin.MIMEJSON, gin.MIMEXML},
			Data:    autos,
		})
		//c.JSON(200, autos)
	}
}

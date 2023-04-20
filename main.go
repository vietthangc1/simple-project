package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code uint32 `json:"code"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/home", func (c *gin.Context) {
		c.JSON(http.StatusFound, ReturnMessage{
			Code: 200,
			Message: "Welcome!",
		})
	})

	router.Run(":80")

}

package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReturnMessage struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusFound, ReturnMessage{
			Code:    200,
			Message: "This is Homepage!",
		})
	})
	router.GET("/home", func(c *gin.Context) {
		c.JSON(http.StatusFound, ReturnMessage{
			Code:    200,
			Message: "Welcome home!",
		})
	})
	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusFound, ReturnMessage{
			Code:    200,
			Message: fmt.Sprintf("Hello %s",name),
		})
	})

	router.Run(":4001")

}

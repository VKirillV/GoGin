package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	variable := "world"
	r := gin.Default()
	r.GET(variable)
	r.GET("/hello/:variable", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Hello": variable,
		})
	})
	r.Run(":8080")
}

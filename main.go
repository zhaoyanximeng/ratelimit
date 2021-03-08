package main

import (
	"github.com/gin-gonic/gin"
	"ratelimit/lib"
)

func test(c *gin.Context) {
	c.JSON(200,gin.H{"message":"ok"})
}

func main() {
	r := gin.New()
	r.GET("/",lib.Limiter(3)(test))

	r.Run(":8080")
}

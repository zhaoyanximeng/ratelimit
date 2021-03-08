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
	//r.GET("/",lib.ParamLimiter(5,1,"name")(lib.Limiter(10)(test)))
	r.GET("/ip_limiter",lib.IpLimiter(10,1)(test))

	r.Run(":8080")
}

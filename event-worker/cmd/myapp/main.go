package main

import (
	"event-worker/internal/handler"

	"github.com/gin-gonic/gin"
)



func main(){
	route := gin.Default()
	route.GET("health", handler.HealthCheck)
	if err := route.Run("8080"); err != nil {
		panic(err)
	}
}
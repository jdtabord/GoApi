package main

import (
	"goApi/request"

	"github.com/gin-gonic/gin"
)

//Execute ejecuta la peticion
func Execute(request request.Request) string {
	if request.Endpoint() == "ping" {
		return "pong"
	}
	return "request unavaliable"
}

func main() {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.Run()
}

func pingHandler(c *gin.Context) {
	get := request.Get{RequestType: "ping"}
	c.JSON(200, gin.H{
		"message": Execute(get),
	})
}

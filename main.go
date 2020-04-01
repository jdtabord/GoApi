package main

import "github.com/gin-gonic/gin"

type Request interface{
	Endpoint() string	
}

type Get struct{
	endpoint string
}

func(this Get) Endpoint() string{	
	return this.endpoint
}

func Execute(request Request) string{
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
	get := Get{"ping"}  	
    c.JSON(200, gin.H{
		"message": Execute(get),
	})
}
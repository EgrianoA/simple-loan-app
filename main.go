package main

import (
	"simple-loan-app/routes"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

var router *gin.Engine

func main() {
	r := gin.Default()
	api := r.Group("/api")
	routes.Routes(api)
	// routes.Routes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}

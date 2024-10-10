package main

import (
	"example.com/rest_api/db"
	"example.com/rest_api/routes"
	"github.com/gin-gonic/gin"
)



func main(){
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)



	server.Run(":8080")
}


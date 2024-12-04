package main

import (
	"udemy-rest-api/db"
	"udemy-rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RefiterRoutes(server)

	server.Run(":8080")
}

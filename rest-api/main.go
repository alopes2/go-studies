package main

import (
	"fmt"

	"github.com/alopes2/go-studies/rest-api/db"
	"github.com/alopes2/go-studies/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080") // localhost

	if err != nil {
		fmt.Println(err)
	}
}

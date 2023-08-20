package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaseelaali/passwordgenerator_unique/database"
	"github.com/jaseelaali/passwordgenerator_unique/routes"
)

func init() {
	database.DatabaseConnection()
}

func main() {
	r := gin.Default()
	routes.User(r)
	r.Run(":8080")
}

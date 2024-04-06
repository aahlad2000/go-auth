package main

import (
	"go-auth/m/controllers"
	"go-auth/m/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DbConnection()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)

	r.POST("/login", controllers.Login)

	r.Run()
}

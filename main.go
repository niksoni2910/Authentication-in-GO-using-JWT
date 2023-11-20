package main

import (
	"go_jwt/controllers"
	"go_jwt/initializers"
	"go_jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadENVvariables()
	initializers.ConnectToDB()
	initializers.SyncDtabase()
}
func main() {
	r := gin.Default()
	r.POST("/signup",controllers.SignUp)
	r.POST("/login",controllers.LogIn)
	r.GET("/validate",middleware.RequireAuth,controllers.Validate)
	r.Run()
}
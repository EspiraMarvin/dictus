package main

import (
	logg "github.com/sirupsen/logrus"

	"github.com/EspiraMarvin/go-crud-postgres/controllers"
	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/EspiraMarvin/go-crud-postgres/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	logg.Info("INIT FUNC RUN!")
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.PATCH("/posts/:id", controllers.PostsUpdate) // partial update
	r.DELETE("/posts/:id", controllers.PostsDelete)

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run() // listen and serve on 0.0.0.0:3000
}

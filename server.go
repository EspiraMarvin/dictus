package main

import (
	logg "github.com/sirupsen/logrus"

	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	logg.Info("INIT FUNC RUN!")
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:3000
}

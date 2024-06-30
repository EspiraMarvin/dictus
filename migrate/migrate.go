package main

import (
	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/EspiraMarvin/go-crud-postgres/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

package controllers

import (
	log "github.com/sirupsen/logrus"

	"github.com/EspiraMarvin/go-crud-postgres/initializers"
	"github.com/EspiraMarvin/go-crud-postgres/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data of req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	log.Info("binded body", &body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return a post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// response with posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// Get a single post
	var post models.Post
	// initializers.DB.First(&post, uuid)
	initializers.DB.Where("uuid = ?", id).First(&post)

	// response with posts
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// get data from req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// find the post we're updating
	var post models.Post
	initializers.DB.Where("uuid = ?", id).First(&post)

	// updating
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})

	// response with updated post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// get id from url
	id := c.Param("id")

	// deleting
	var post models.Post
	// initializers.DB.Delete(&post, id)
	initializers.DB.Where("uuid = ?", id).Delete(&post)

	// reponse
	c.Status(200)
}

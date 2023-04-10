package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// We need to get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Also need to create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Then need to return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	// Get ID from URL
	id := c.Param("id")

	// Get the posts
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
})
}

func PostsUpdate(c *gin.Context) {
	// Get ID off URL

	id := c.Param("id")


	// Get data off req body

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find post we are updating
	var post models.Post
	initializers.DB.First(&post, id)


	// Update it

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title, 
		Body: body.Body,
	})

	// Respond with it

	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsDelete(c *gin.Context) {
	// Get ID from URL
	id := c.Param("id")

	// Delete Posts

	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.Status(200)
}
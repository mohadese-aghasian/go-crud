package controllers

import (
	"example.com/go/config"
	"example.com/go/models"
	"github.com/gin-gonic/gin"
)

func BlogCreate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	blog := models.Blog{Title: body.Title, Body: body.Body}

	result := config.DB.Create(&blog)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"blog": blog,
	})

}

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "finally",
	})
}

func BlogList(c *gin.Context) {

	var blogs []models.Blog
	config.DB.Find(&blogs)

	c.JSON(200, gin.H{
		"blogs": blogs,
	})
}

func GetBlog(c *gin.Context) {

	id := c.Param("id")

	var blog models.Blog
	config.DB.First(&blog, id)

	c.JSON(200, gin.H{
		"blogs": blog,
	})
}

func UpdateBlog(c *gin.Context) {

	var body struct {
		ID    uint
		Title string
		Body  string
	}

	c.Bind(&body)

	var blog models.Blog
	config.DB.First(&blog, body.ID)

	// blog := models.Blog{ID: body.ID, Title: body.Title, Body: body.Body}

	blog.Title = body.Title
	blog.Body = body.Body

	config.DB.Save(&blog)

	c.JSON(200, gin.H{
		"blogs": blog,
	})
}

func DeleteBlog(c *gin.Context) {

	id := c.Param("id")

	// var blog models.Blog
	// config.DB.First(&blog, id)

	config.DB.Delete(&Blog, id)

	c.JSON(200, gin.H{
		"blogs": blog,
	})
}

package main

import (
	"example.com/go/config"
	"example.com/go/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	// config.LoadEnvVariables()
	config.ConnectToDB()

}

func main() {

	// err := migrate.mig()

	// if err != nil {
	// 	panic("Migration failed: " + err.Error())
	// }
	// fmt.Println("Migration completed!")

	r := gin.Default()
	r.GET("/", controllers.Index)
	r.POST("/blogs", controllers.BlogCreate)
	r.GET("/blogs", controllers.BlogList)
	r.GET("/blogs/:id", controllers.GetBlog)
	r.PUT("/blogs", controllers.UpdateBlog)
	r.DELETE("/blogs/:id", controllers.DeleteBlog)

	r.Run(":3000")
}

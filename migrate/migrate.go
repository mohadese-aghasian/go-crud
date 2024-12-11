package main

import (

	"example.com/go/config"
	"example.com/go/models"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectToDB()
	
}

func main() {
	config.DB.AutoMigrate(&models.Blog{})
	// err := config.DB.AutoMigrate(&models.Blog{})
	
	// if err != nil {
	// 	panic("Failed to migrate database: " + err.Error())
	// }
	// fmt.Println("Database migration completed!")
}

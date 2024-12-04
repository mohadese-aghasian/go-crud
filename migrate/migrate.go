package main

import (
	"go/initializers"

	"go/models"
)

func init() {
	initializers.ConnectToDB()
}

func main(){
	initializers.db.AutoMigrate(&models.Blog{})
}
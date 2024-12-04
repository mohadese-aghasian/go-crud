package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)	

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres76555432"
	dbName   = "testGoDb"
)
func ConnectToDB()*gorm.DB{
	
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbName, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!=nil {
		log.Fatal("Failed database connection!")
	}

	fmt.Println("Database connection successful!")
	return db

}
package main

import (
	"fmt"
	"log"
	"registration/model"
	"registration/routes"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	dsn := "root:@tcp(127.0.0.1:3306)/registration?charset=utf8mb4&parseTime=True"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed")
	}

	_ = db

	fmt.Println("DB connected")
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Role{})
	// db.AutoMigrate(&model.UserRole{})
	log.Println("Database migration completed")

	

	apiRouter := routes.ApiRoute(db)
	apiRouter.Run(":8090")
}
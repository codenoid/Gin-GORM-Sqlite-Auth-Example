package main

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// init database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	mainDB = db

	// Migrate the schema
	if err := db.AutoMigrate(&UserModel{}); err != nil {
		log.Println(err)
	}

	hashedPassword, _ := HashPassword("test123")
	db.Create(&UserModel{Name: "Rubi", Username: "rubi", Password: hashedPassword})

	// init go cache
	goCache = cache.New(5*time.Minute, 10*time.Minute)

}

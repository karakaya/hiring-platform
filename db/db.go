package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB{
	db,err:=gorm.Open(sqlite.Open("hiring.db"),&gorm.Config{})
	if err != nil{
		log.Println(err)
	}
	return db
}

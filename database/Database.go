package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)
var DB *gorm.DB
func InitDB(){
	db,err:=gorm.Open(sqlite.Open("hiring.db"),&gorm.Config{})
	if err != nil{
		log.Println(err)
	}
	DB = db
}

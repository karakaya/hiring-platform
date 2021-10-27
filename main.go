package main

import (
	"github.com/gorilla/mux"
	"hiring-platform/company"
	"hiring-platform/db"
	"log"
	"net/http"
)

func main(){
	r:= mux.NewRouter()
 	company.Route(r)

	err := db.InitDB().AutoMigrate(&db.Company{},&db.Hr{},&db.JobAdvert{}); if err != nil{
		log.Printf("err to migrate db: %v",err)
	}

	err = http.ListenAndServe(":8080",r)
	if err != nil{
	 log.Println(err)
 	}

}
package main

import (
	"github.com/gorilla/mux"
	"hiring-platform/company"
	"hiring-platform/database"
	"hiring-platform/hr"
	"log"
	"net/http"
)

func main(){
	r:= mux.NewRouter()
	hr.Route(r)
 	company.Route(r)
	database.InitDB()
	err := database.DB.AutoMigrate(&database.Company{},&database.Hr{},&database.JobAdvert{},&database.InviteHr{}); if err != nil{
		log.Printf("err to migrate db: %v",err)
	}
	
	err = http.ListenAndServe(":8080",r)
	if err != nil{
	 log.Println(err)
 	}

}
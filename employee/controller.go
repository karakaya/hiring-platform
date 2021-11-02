package employee

import (
	"encoding/json"
	"hiring-platform/database"
	"log"
	"net/http"
)

func apply(w http.ResponseWriter, r *http.Request){
	var apply database.Applicants
	err := json.NewDecoder(r.Body).Decode(&apply)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	database.DB.Create(&apply)
	w.WriteHeader(http.StatusCreated)
}

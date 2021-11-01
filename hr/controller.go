package hr

import (
	"encoding/json"
	"fmt"
	"hiring-platform/database"
	"log"
	"net/http"
)

func join(w http.ResponseWriter, r *http.Request) {
	//TODO add self expiring link
	var join map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&join)
	if err != nil {
		log.Println(err)
	}

	var hr database.Hr
	database.DB.Model(&database.Hr{}).Where("id", join["hr_id"]).Scan(&hr)

	var invite database.InviteHr
	database.DB.Model(&database.InviteHr{}).Where("hr_id",join["hr_id"]).Scan(&invite)

	hr.CompanyID = invite.CompanyID
	fmt.Println(hr)
	database.DB.Save(&hr)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("successfully joined to the company"))

}
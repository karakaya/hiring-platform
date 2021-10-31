package company

import (
	"encoding/json"
	"github.com/google/uuid"
	"hiring-platform/database"
	amqp "hiring-platform/rabbitmq"
	"log"
	"net/http"
	"time"
)

func index(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("controller"))
	if err != nil {
		return 
	}
}

func register(w http.ResponseWriter, r *http.Request){
	var company database.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	failOnError(err,"err to decode company")
	database.DB.Create(&company)
}
//invite your hr to team to your company
func invite(w http.ResponseWriter, r *http.Request){
	var hr database.Hr
	err:=json.NewDecoder(r.Body).Decode(&hr)
	failOnError(err,"err to decode hr")

	database.DB.Create(&hr)

	failOnError(err,"err to marshal json")
	invite := database.InviteHr{Name: hr.Name,Email: hr.Email,CompanyID: hr.CompanyID,HrID: hr.ID,CreatedAt: time.Now(),Link: uuid.New()}
	database.DB.Create(&invite)
	data,err := json.Marshal(invite)
	amqp.Produce(data)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
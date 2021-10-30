package company

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hiring-platform/db"
	amqp "hiring-platform/rabbitmq"
	"log"
	"net/http"
	"time"
)
var dbc *gorm.DB
func init(){
	dbc = db.InitDB()
}
func index(w http.ResponseWriter, r *http.Request){
	_, err := w.Write([]byte("controller"))
	if err != nil {
		return 
	}
}



func register(w http.ResponseWriter, r *http.Request){
	var company db.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	failOnError(err,"err to decode company")
	dbc.Create(&company)
}
//invite your hr to team to your company
func invite(w http.ResponseWriter, r *http.Request){
	var hr db.Hr
	err:=json.NewDecoder(r.Body).Decode(&hr)
	failOnError(err,"err to decode hr")
	dbc.Create(&hr)

	failOnError(err,"err to marshal json")
	invite := db.InviteHr{Name:hr.Name,Email: hr.Email,CompanyID: hr.CompanyID,HrID: hr.ID,CreatedAt: time.Now(),Link: uuid.New()}
	dbc.Create(&invite)
	data,err := json.Marshal(invite)

	amqp.Produce(data)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
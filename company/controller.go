package company

import (
	"encoding/json"
	"gorm.io/gorm"
	"hiring-platform/db"
	amqp "hiring-platform/rabbitmq"
	"log"
	"net/http"
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
	if err != nil{
		log.Printf("err decode create company form: %v ",err)
	}
	dbc.Create(&company)
}
//invite your hr to team to your company
func invite(w http.ResponseWriter, r *http.Request){
	var hr db.Hr
	json.NewDecoder(r.Body).Decode(&hr)
	db.InitDB().Create(&hr)
	data,err := json.Marshal(hr)
	if err != nil{
		log.Printf("err to marshal hr: %v",err)
	}
	amqp.Produce(data)
}
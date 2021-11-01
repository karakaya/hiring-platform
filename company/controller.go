package company

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"hiring-platform/database"
	amqp "hiring-platform/rabbitmq"
	"log"
	"net/http"
	"time"
)



func register(w http.ResponseWriter, r *http.Request){
	var company database.Company
	err := json.NewDecoder(r.Body).Decode(&company)
	hash,_ := HashPassword(company.Password)
	company.Password = hash
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

func jobAdvert(w http.ResponseWriter, r *http.Request){
	//TODO company id from auth token
	var advert database.JobAdvert

	err := json.NewDecoder(r.Body).Decode(&advert)
	failOnError(err,"err to decode advert body")
	fmt.Println(advert)

	database.DB.Create(&advert)
	w.WriteHeader(http.StatusCreated)
	
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
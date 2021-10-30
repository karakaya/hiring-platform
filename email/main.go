package email

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	gomail "gopkg.in/mail.v2"
	"hiring-platform/db"
	"log"
)

func Send(data []byte){

	var invite db.InviteHr
	err := json.Unmarshal(data,&invite)
	if err != nil{
		log.Println(err)
		return
	}

	m:= gomail.NewMessage()

	m.SetHeader("From","from@here.com")
	m.SetHeader("To",invite.Email)
	m.SetHeader("Subject","Hi, join to the hiring-platform")

	bodyToSend :=fmt.Sprintf("hi %s, your comapny invited you to join hiring-platform. here is the <a href=localhost:8080/hr/join/\"%s\">link</a>, click and join.",invite.Name,invite.Link)

	m.SetBody("text/html",bodyToSend)

	d:=gomail.NewDialer("smtp.mailtrap.io",25,"","")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil{
		fmt.Println(err)
		panic(err)
	}
}


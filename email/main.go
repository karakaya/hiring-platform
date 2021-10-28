package email

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/mail.v2"
)

func Send(data []byte){
	m:= gomail.NewMessage()
	m.SetHeader("From","from@here.com")
	m.SetHeader("To","to@example.com")
	m.SetHeader("Subject","subject here")
	m.SetBody("text/plain","here is the body")
	d:=gomail.NewDialer("smtp.mailtrap.io",25,"","")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil{
		fmt.Println(err)
		panic(err)
	}
	return
}


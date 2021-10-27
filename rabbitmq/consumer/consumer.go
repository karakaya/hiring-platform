package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main(){
	conn,err:=amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil{
		log.Printf("err to dial rabbitmq: %v",err)
	}
	defer conn.Close()

	ch,err:=conn.Channel()
	if err != nil{
		log.Printf("err to create channel: %v",err)
	}
	defer ch.Close()




	messages,err := ch.Consume(
		"invite-hr",
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil{
		log.Printf("err to declare qeue: %v",err)
	}
	log.Println("waiting for rabbitmq messages")
	forever := make(chan bool)

	go func(){
		for message := range messages{
			log.Printf("> recieved message: %s \n",message.Body)
		}
	}()
	<-forever
}
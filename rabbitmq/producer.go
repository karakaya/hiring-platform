package rabbitmq

import (
	"github.com/streadway/amqp"

	"log"
)

func Produce(data []byte){
	conn,err:= amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil{
		log.Printf("err to connect to rabbitmq: %v",err)
	}
	ch,err:=conn.Channel()
	if err != nil{
		log.Printf("err connect to channel: %v",err)
	}

	defer ch.Close()
	defer conn.Close()

	_,err=ch.QueueDeclare(
		"invite-hr2",
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil{
		log.Printf("err to declare queue: %v",err)
	}
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body: data,
	}
	if err = ch.Publish(
		"",
		"invite-hr",
		false,
		false,
		message,
		); err != nil{
		log.Printf("err publishing the message: %v",err)
	}
}


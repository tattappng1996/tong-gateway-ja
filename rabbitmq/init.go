package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var RabbitConn *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func RabbitMQInit() {
	fmt.Println("rabbitmq is connecting..")

	var err error
	RabbitConn, err = amqp.Dial("amqp://user:VimRkNp1VjGJ@35.186.158.107:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	fmt.Println("rabbitmq connection success!")
}

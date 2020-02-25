package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func Server(exchangeName string, directKey string, queueName string, function func(amqp.Delivery) []byte) {
	channel, err := RabbitConn.Channel()
	failOnError(err, "Failed to open a channel")

	err = channel.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	queue, err := channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = channel.QueueBind(
		queue.Name,   // queue name
		directKey,    // routing key
		exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")

	err = channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	messages, err := channel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for delivery := range messages {
			go func(delivery amqp.Delivery) {
				response := function(delivery)
				err = channel.Publish(
					"",               // exchange
					delivery.ReplyTo, // routing key
					false,            // mandatory
					false,            // immediate
					amqp.Publishing{
						ContentType:   "application/json",
						CorrelationId: delivery.CorrelationId,
						Body:          response,
					})
				failOnError(err, "Failed to publish")
			}(delivery)
		}
	}()

	log.Printf(" [server] running exchange=%s, directKey=%s, queue=%s", exchangeName, directKey, queueName)
	<-forever
}

func Client(messages []byte, exchangeName, directKey string, corrID string) []byte {
	channel, err := RabbitConn.Channel()
	failOnError(err, "Failed to open a channel")

	err = channel.ExchangeDeclare(
		exchangeName, // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a exchange")

	consumer, err := channel.Consume(
		"amq.rabbitmq.reply-to", // queue
		"",                      // consumer
		true,                    // auto-ack
		false,                   // exclusive
		false,                   // no-local
		false,                   // no-waitx``
		nil,                     // args
	)
	failOnError(err, "Failed to publish a message")

	err = channel.Publish(
		exchangeName, // exchange
		directKey,    // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType:   "text/plain",
			CorrelationId: corrID,
			ReplyTo:       "amq.rabbitmq.reply-to",
			Body:          messages,
		})
	failOnError(err, "Failed to publish a message")

	// fmt.Printf(" [x] Request: %s\n", message)

	for delivery := range consumer {
		if corrID == delivery.CorrelationId {

			return delivery.Body
		}
		// log.Printf(" [x] %s", d.Body)
	}

	// log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	return nil
}

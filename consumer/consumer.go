package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, _ := amqp.Dial("amqp://root:root@localhost:16200")
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"firstExample",
		true,
		false,
		false,
		false,
		nil)

	msgs, _ := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	for m := range msgs {
		fmt.Println(string(m.Body))
	}
}

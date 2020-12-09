package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	uri string = string("amqp://root:root@localhost:16200")
)

func main() {
	fmt.Println(uri)
	conn, err := amqp.Dial(uri)

	if err != nil {
		log.Fatalf("Dial: %s", err)
	}
	publish("test message", conn)
}

func publish(message string, conn *amqp.Connection) {
	ch, _ := conn.Channel()

	q, _ := ch.QueueDeclare(
		"firstExample",
		true,  //durable
		false, //autodelete
		false, //exclusive
		false, //noWait
		nil)

	msg := amqp.Publishing{
		Body: []byte("example go")}

	ch.Publish(
		"",
		q.Name,
		false,
		false,
		msg)
}

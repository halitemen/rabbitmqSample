package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

var (
	uri string = string("amqp://root:root@localhost:16200")
)

func main() {
	conn, err := amqp.Dial(uri)

	if err != nil {
		log.Fatalf("Dial: %s", err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Message: ")
	msgUser, _ := reader.ReadString('\n')

	publish(msgUser, conn)
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
		Body: []byte(message)}

	err := ch.Publish(
		"",
		q.Name,
		false,
		false,
		msg)

	if err != nil {
		log.Fatalf("Dial: %s", err)
	}
}

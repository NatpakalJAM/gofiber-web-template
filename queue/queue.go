package queue

import (
	"encoding/json"
	"gofiber-web-template/cfg"
	"gofiber-web-template/model"
	"log"
	"time"

	"fmt"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel
var queueProcess amqp.Queue

//Init init queue instance
func Init() {
	var err error
	qCfg := cfg.C.RabbitMQ
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%v/",
		qCfg.Username, qCfg.Password, qCfg.Host, qCfg.Port)
	conn, err = amqp.Dial(connStr)
	if err != nil {
		log.Fatal(err)
	}

	ch, err = conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	queueProcess, err = ch.QueueDeclare(
		fmt.Sprintf("%s-%s", qCfg.Queues.Process, cfg.C.Instance), // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

//PubishProcessQueue publish to process queue
func PubishProcessQueue(ID uint64, data interface{}) (err error) {
	msg := model.QueueMessage{
		ID:   ID,
		Data: data,
	}
	body, err := json.Marshal(&msg)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"",                // exchange
		queueProcess.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		})

	if err != nil {
		return err
	}

	return nil
}

//StartProcessConsumer start process consumer
func StartProcessConsumer() {
	msgs, err := ch.Consume(
		queueProcess.Name, // queue
		"",                // random-consumer-name
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			msg := model.QueueMessage{}
			err := json.Unmarshal(d.Body, &msg)
			if err != nil {
				log.Printf("Error unmarshal message: %s", err.Error())
				continue
			}
			/*json unmarshal
			then switch process by activity id
			*/
		}
	}()

	<-forever
}

package agcmq

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessageQueue(argQueueName string, argBody string) {

	conn := agcmq.NewConnection(argQueueName)
	
	if err := conn.Connect(); err != nil {
		panic(err)
	}

	defer conn.ChannelClose()
	if err := conn.Publish(argBody); err != nil {
		panic(err)
	}
}
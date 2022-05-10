package agcmq

import (
	"fmt"
	
	//amqp "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
)

//Consume consumes the messages from the queues and passes it as map of chan of amqp.Delivery
func (c *Connection) Consume() (<-chan amqp.Delivery, error) {
	err := c.channel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to set QoS: %s", err)
	}

	msgs, err := c.channel.Consume(
		c.queue_name, // queue
		"",           // consumer
		false,        // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to register a consumer : %s", err)
	}
	return msgs, nil

}
package agcmq

import (
	"fmt"
	"log"

	//amqp "github.com/rabbitmq/amqp091-go"
	"github.com/streadway/amqp"
)

//Publish publishes a request to the amqp queue
func (c *Connection) Publish(body string) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	err := c.channel.Publish(
		"",           // exchange
		c.queue_name, // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Publishing: %s", err)
	}
	log.Printf(" [x] Sent %s", body)
	return nil
}

func (c *Connection) PublishBackup(body []byte) error {
	select { //non blocking channel - if there is no error will go to default where we do nothing
	case err := <-c.err:
		if err != nil {
			c.Reconnect()
		}
	default:
	}

	err := c.channel.Publish(
		"",           // exchange
		c.queue_name, // routing key
		false,        // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		return fmt.Errorf("Publishing: %s", err)
	}
	log.Printf(" [x] Sent %s", body)
	return nil
}
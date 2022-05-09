package agcmq

import (
	"errors"
	"fmt"
	
	// "github.com/streadway/amqp"
	amqp "github.com/rabbitmq/amqp091-go"	
)


const amqpURI = "amqp://admin:Ketilinux1@10.0.7.4:31011/"
//const amqpURI = "amqp://admin:Ketilinux1@mq:5672/"


const (
	MQ_DockerfileMaker = "MAKER_WAIT_QUEUE"
	MQ_DockerfileBuilder = "BUILD_WAIT_QUEUE"
	MQ_InferenceWait = "INFERENCE_WAIT_QUEUE"
	MQ_QueryCommander = "QUERYING_WAIT_QUEUE"

	MQ_NewAGC_TestQueue = "AGC_TEST_QUEUE"
)

//Connection is the connection created
type Connection struct {
	queue_name string
	conn       *amqp.Connection
	channel    *amqp.Channel
	err        chan error
}

var (
	connectionPool = make(map[string]*Connection)
)

//NewConnection returns the new connection object
func NewConnection(queueName string) *Connection {
	if c, ok := connectionPool[queueName]; ok {
		return c
	}
	c := &Connection{
		queue_name: queueName,
		err:        make(chan error),
	}
	connectionPool[queueName] = c
	return c
}

func (c *Connection) Connect() error {
	var err error

	c.conn, err = amqp.Dial(amqpURI)
	if err != nil {
		return fmt.Errorf("Error in creating rabbitmq connection with %s : %s", amqpURI, err.Error())
	}
	go func() {
		<-c.conn.NotifyClose(make(chan *amqp.Error)) //Listen to NotifyClose
		c.err <- errors.New("Connection Closed")
	}()
	c.channel, err = c.conn.Channel()
	if err != nil {
		return fmt.Errorf("Channel: %s", err)
	}

	_, err = c.channel.QueueDeclare(
		c.queue_name, // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return fmt.Errorf("QueueDeclare: %s", err)
	}
	return nil
}

func (c *Connection) ChannelClose() {
	defer c.channel.Close()
}

//Reconnect reconnects the connection
func (c *Connection) Reconnect() error {
	if err := c.Connect(); err != nil {
		return err
	}
	return nil
}
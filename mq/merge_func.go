package agcmq


func SendMessageQueue(argQueueName string, argBody string) {
	conn := NewConnection(argQueueName)
	
	if err := conn.Connect(); err != nil {
		panic(err)
	}

	defer conn.ChannelClose()
	if err := conn.Publish(argBody); err != nil {
		panic(err)
	}
}
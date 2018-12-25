package kafka

import (
	"fmt"
	"github.com/shopify/sarama"
	"xlog"
)

var (
	client sarama.SyncProducer
	msgChan chan *Message
)

type Message struct {
	Line string
	Topic string
}

func Init(address []string,chanSize int) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		xlog.LogError("producer close, err:", err)
		return
	}
	msgChan = make(chan *Message,chanSize)
	go sendKafka()
	return
}

func sendKafka() {
	for msg := range msgChan {
		kafkaMsg := &sarama.ProducerMessage{}
		kafkaMsg.Topic = msg.Topic
		kafkaMsg.Value = sarama.StringEncoder(msg.Line)

		pid, offset, err := client.SendMessage(kafkaMsg)
		if err != nil {
			xlog.LogError("Send message failed,",err)
			continue
		}
		xlog.LogDebug("PID:%v offset:%v\n",pid,offset )
	}
}

func SendLog(msg *Message) (err error) {
	if len(msg.Line) == 0 {
		return
	}
	select {
	case msgChan <- msg:
	default:
		err = fmt.Errorf("Chan is full")
	}
	return
}

package kafka_client

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

type KafkaMessage struct {
	Key   []byte
	Value []byte
}

type KafkaProducer struct {
	writer     *kafka.Writer
	messageCh  chan KafkaMessage
	workerDone chan struct{}
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	kp := &KafkaProducer{
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Topic:    topic,
			Balancer: &kafka.LeastBytes{},
		},
		messageCh:  make(chan KafkaMessage, 100),
		workerDone: make(chan struct{}, 1),
	}

	go kp.startWorker()

	return kp
}

func (kp *KafkaProducer) startWorker() {
	for msg := range kp.messageCh {
		err := kp.writer.WriteMessages(context.TODO(), kafka.Message{
			Key:   msg.Key,
			Value: msg.Value,
		})

		if err != nil {
			log.Printf("failed to send kafka message: %v", err)
		}
	}
	kp.workerDone <- struct{}{}
}

func (kp *KafkaProducer) SendMessage(msg KafkaMessage) {
	kp.messageCh <- msg
}

func (kp *KafkaProducer) Close() error {
	close(kp.messageCh)

	for i := 0; i < cap(kp.workerDone); i++ {
		<-kp.workerDone
	}

	return kp.writer.Close()
}

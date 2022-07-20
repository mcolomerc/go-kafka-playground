package producer

import (
	"mcolomerc/kafka-playground/config"
)

type ProducerArgs struct {
	TopicPrefix string
	TopicNumber int
	Messages    int
	Size        int
	Keys        []string
	Values      []string
}

type Producer interface {
	Produce(args ProducerArgs) (ProducerResponse, error)
	Close()
}

type KafkaProducer struct {
	Cfg  *config.Config
	Name string
}

type ProducerResponse struct {
	Producer       string   `json:"producer"`
	Topics         []string `json:"topics"`
	NumberMessages int      `json:"number_messages"`
	SizeMessage    int      `json:"size_message"`
	Duration       string   `json:"duration"`
}

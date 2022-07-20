package confluent

import (
	"encoding/json"
	"fmt"
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/producer"
	"mcolomerc/kafka-playground/utils"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/prometheus/client_golang/prometheus"
)

type ConfluentProducer struct {
	producer.KafkaProducer
	ProducerConfig *kafka.ConfigMap
	Producer       *kafka.Producer
}

func New(cfg config.Config, name string) (*ConfluentProducer, error) {
	producerConfig := &kafka.ConfigMap{} // create config map for producer
	for k, v := range cfg.KAFKA_CONFIGS {
		producerConfig.SetKey(k, v)
	}

	// Prometheus collector
	statsCollector := NewStatsCollector()
	prometheus.MustRegister(statsCollector)

	return &ConfluentProducer{
		KafkaProducer: producer.KafkaProducer{
			Cfg:  &cfg,
			Name: name,
		},
		ProducerConfig: producerConfig,
	}, nil
}

func (p *ConfluentProducer) Produce(args producer.ProducerArgs) (producer.ProducerResponse, error) {
	resp := producer.ProducerResponse{
		Topics:         []string{},
		NumberMessages: 0,
		SizeMessage:    args.Size,
		Duration:       "",
	}
	msgcnt := 0
	start := time.Now()
	send := make(chan int, args.Messages)
	prod, err := kafka.NewProducer(p.ProducerConfig)
	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		return resp, err
	}
	p.Producer = prod
	// Listen to all the events on the default events channel
	go func() {
		for e := range prod.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				m := ev
				if m.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
				} else {
					/* fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
					*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset) */
					send <- 1
				}
			case kafka.Error:
				fmt.Printf("Error: %v\n", ev)
			case *kafka.Stats: // stats event
				var statsResult Stats
				json.Unmarshal([]byte(e.String()), &statsResult)
				ProcessStats(statsResult)

			default:
				fmt.Printf("Ignored event: %s\n", ev)
			}
		}
	}()
	topics := []string{}
	//Writing messages to Kafka
	for msgcnt < args.Messages {
		topic := args.TopicPrefix + "_" + p.Name + "_" + strconv.Itoa(msgcnt%args.TopicNumber)
		topics = append(topics, topic)
		err := p.Producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Key:            []byte(args.Keys[msgcnt]),
			Value:          []byte(args.Values[msgcnt]),
			Headers:        nil,
		}, nil)

		if err != nil {
			if err.(kafka.Error).Code() == kafka.ErrQueueFull {
				// Producer queue is full, wait 1s for messages
				// to be delivered then try again.
				time.Sleep(time.Second)
				continue
			}
			fmt.Printf("Failed to produce message: %v\n", err)
		}
		msgcnt++
	}

	elapsed := time.Since(start)
	sent := 0
	for i := 0; i < args.Messages; i++ {
		<-send
		sent++
	}

	resp.Duration = elapsed.String()
	resp.Topics = utils.UniqueNonEmptyElementsOf(topics)
	resp.NumberMessages = sent
	resp.Producer = p.Name
	log.Println(resp)
	return resp, nil
}

func (p *ConfluentProducer) Close() {
	p.Producer.Close()
}

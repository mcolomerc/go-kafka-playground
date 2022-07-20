package sarama

import (
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/producer"
	"mcolomerc/kafka-playground/utils"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

type SaramaProducer struct {
	producer.KafkaProducer
	SaramaConfig   *sarama.Config
	Brokers        []string
	SaramaProducer sarama.AsyncProducer
}

func New(cfg config.Config, name string) (*SaramaProducer, error) {
	// setupProducer will create a AsyncProducer and returns it
	kafkaBrokers := []string{}
	config := sarama.NewConfig()
	config.Metadata.AllowAutoTopicCreation = true
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.Producer.MaxMessageBytes = 1000000
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches
	// config.Producer.Flush.Bytes = *flushBytes
	// config.Producer.Flush.Messages = *flushMessages
	// config.Producer.Flush.MaxMessages = *flushMaxMessages
	// KAFKA_BOOTSTRAP_SERVERS
	for k, v := range cfg.KAFKA_CONFIGS {
		if k == "bootstrap.servers" {
			kafkaBrokers = append(kafkaBrokers, v.(string))
		}
		if k == "batch.size" {
			bSize, _ := strconv.Atoi(v.(string))
			config.Producer.MaxMessageBytes = bSize
		}
	}

	sp, err := sarama.NewAsyncProducer(kafkaBrokers, config)
	if err != nil {
		log.Println("Failed to start Sarama producer:", err)
		return nil, err
	}

	return &SaramaProducer{
		KafkaProducer: producer.KafkaProducer{
			Cfg:  &cfg,
			Name: name,
		},
		SaramaConfig:   config,
		SaramaProducer: sp,
	}, nil
}

func (p *SaramaProducer) Produce(args producer.ProducerArgs) (producer.ProducerResponse, error) {

	resp := producer.ProducerResponse{
		Topics:         []string{},
		NumberMessages: 0,
		SizeMessage:    args.Size,
		Duration:       "",
	}
	send := make(chan int, args.Messages)

	msgcnt := 0
	start := time.Now()

	topics := []string{}
	//Writing messages to Kafka
	for msgcnt < args.Messages {
		topic := args.TopicPrefix + "_" + p.Name + "_" + strconv.Itoa(msgcnt%args.TopicNumber)
		topics = append(topics, topic)

		// be distributed randomly over the different partitions.
		message := &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.StringEncoder(args.Keys[msgcnt]),
			Value: sarama.StringEncoder(args.Values[msgcnt]),
		}
		// Asynchronous sending only returns after writing to memory, but it is not really sent out
		// The sarama library uses a channel to receive messages. The background goroutine asynchronously takes messages from the channel and sends them
		p.SaramaProducer.Input() <- message

		go func() {
			// [! important] after sending, the asynchronous producer must read the return value from Errors or Successes.
			// Otherwise, the internal processing logic of sarama will be blocked and only one message can be sent
			select {
			case _ = <-p.SaramaProducer.Successes():
				send <- 1
			case e := <-p.SaramaProducer.Errors():
				if e != nil {
					log.Printf("[Producer] err:%v msg:%+v \n", e.Msg, e.Err)
				}
			}
		}()
		msgcnt++

	}

	log.Println(msgcnt)
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

func (p *SaramaProducer) Close() {
	return
}

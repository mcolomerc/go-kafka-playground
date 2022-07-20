package franz

import (
	"context"
	"fmt"
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/producer"
	"mcolomerc/kafka-playground/utils"

	"strconv"
	"sync/atomic"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
)

type FranzProducer struct {
	producer.KafkaProducer
	Client *kgo.Client
	Config []kgo.Opt
}

func New(cfg config.Config, name string) (*FranzProducer, error) {
	seeds := []string{}
	var batchSize int32
	batchSize = 16384
	// KAFKA_BOOTSTRAP_SERVERS
	for k, v := range cfg.KAFKA_CONFIGS {
		if k == "bootstrap.servers" {
			seeds = append(seeds, v.(string))
		}
		if k == "batch.size" {
			bSize, _ := strconv.ParseInt(v.(string), 10, 32)
			batchSize = int32(bSize)
		}
	}

	// metrics := kprom.NewMetrics("franz")

	opts := []kgo.Opt{
		kgo.SeedBrokers(seeds...),
		kgo.AllowAutoTopicCreation(),
		kgo.ProducerBatchMaxBytes(batchSize),
		//	kgo.WithHooks(metrics),
	}

	return &FranzProducer{
		KafkaProducer: producer.KafkaProducer{
			Cfg:  &cfg,
			Name: name,
		},
		Config: opts,
	}, nil
}

var rateRecs int64
var rateBytes int64

func (p *FranzProducer) Produce(args producer.ProducerArgs) (producer.ProducerResponse, error) {
	resp := producer.ProducerResponse{
		Topics:         []string{},
		NumberMessages: 0,
		SizeMessage:    args.Size,
		Duration:       "",
	}
	// One client can both produce and consume!
	client, err := kgo.NewClient(p.Config...)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	p.Client = client
	if p.CheckConnection() == false {
		return resp, fmt.Errorf("Kafka connection is not available")
	}

	msgcnt := 0
	start := time.Now()
	send := make(chan int, args.Messages)

	ctx := context.Background()

	// Calling NewTicker method
	Ticker := time.NewTicker(1 * time.Second)

	// Creating channel using make
	// keyword
	tickerChannel := make(chan bool)

	// Go function
	go func() {
		// Using for loop
		for {
			// Select statement
			select {
			// Case statement
			case <-tickerChannel:
				return
			// Case to print current time
			case <-Ticker.C:
				recs := atomic.SwapInt64(&rateRecs, 0)
				bytes := atomic.SwapInt64(&rateBytes, 0)
				log.Printf(">> %0.2f MiB/s; %0.2fk records/s\n", float64(bytes)/(1024*1024), float64(recs)/1000)
			}
		}
	}()
	topics := []string{}
	//Writing messages to Kafka
	for msgcnt < args.Messages {
		topic := args.TopicPrefix + "_" + p.Name + "_" + strconv.Itoa(msgcnt%args.TopicNumber)
		topics = append(topics, topic)
		record := &kgo.Record{Topic: topic, Key: []byte(args.Keys[msgcnt]), Value: []byte(args.Values[msgcnt])}

		p.Client.Produce(ctx, record, func(record *kgo.Record, err error) {
			if err != nil {
				log.Printf("record had a produce error: %v\n", err)
			} else {
				send <- 1
				atomic.AddInt64(&rateRecs, 1)
				atomic.AddInt64(&rateBytes, int64(len(record.Value)))
			}
		})
		msgcnt++
	}

	elapsed := time.Since(start)
	sent := 0
	for i := 0; i < args.Messages; i++ {
		<-send
		sent++
	}

	// Calling Stop() method
	Ticker.Stop()

	// Setting the value of channel
	tickerChannel <- true

	resp.Duration = elapsed.String()
	resp.Topics = utils.UniqueNonEmptyElementsOf(topics)
	resp.NumberMessages = sent
	resp.Producer = p.Name
	log.Println(resp)
	return resp, nil
}

func (p *FranzProducer) CheckConnection() bool {
	//Check connection
	ctx := context.Background()
	e := p.Client.Ping(ctx)
	if e != nil {
		log.Fatal(e)
		return false
	}
	return true
}

func (p *FranzProducer) Close() {
	p.Client.Close()
}

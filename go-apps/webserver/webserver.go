package webserver

import (
	"encoding/json"
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/data"
	"mcolomerc/kafka-playground/producer"
	"mcolomerc/kafka-playground/utils"

	"net/http"
	"time"
)

type WebHandler struct {
	Config    config.Config
	Data      *data.MockData
	Producers []producer.Producer
}

type Params struct {
	TopicPrefix string `json:"topic_prefix,omitempty"`
	TopicNumber int    `json:"topic_number,omitempty"`
	Messages    int    `json:"messages,omitempty"`
	Size        int    `json:"size,omitempty"`
}

func New(config config.Config, producers []producer.Producer) *WebHandler {
	mData := data.New(config)

	return &WebHandler{
		Config:    config,
		Producers: producers,
		Data:      mData,
	}
}

func (h *WebHandler) Produce(w http.ResponseWriter, r *http.Request) {
	defer utils.TimeTrack(time.Now(), "Produce request")

	p := Params{
		TopicPrefix: h.Config.TOPIC_PREFIX,
		TopicNumber: h.Config.NB_TOPICS,
		Messages:    h.Config.NB_MESSAGES,
		Size:        h.Config.MESSAGE_SIZE,
	}
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	keys, values := h.Data.BuildData(p.Messages, p.Size)
	producerArgs := producer.ProducerArgs{
		TopicPrefix: p.TopicPrefix,
		TopicNumber: p.TopicNumber,
		Messages:    p.Messages,
		Size:        p.Size,
		Keys:        keys,
		Values:      values,
	}

	done := make(chan producer.ProducerResponse, len(h.Producers))
	for _, prod := range h.Producers {
		go func(p *Params, prod producer.Producer) {
			resp, err := prod.Produce(producerArgs)
			if err != nil {
				log.Printf("Error: %s", err)
			}
			done <- resp
		}(&p, prod)
	}
	var responses []producer.ProducerResponse
	for i := 0; i < len(h.Producers); i++ {
		responses = append(responses, <-done)
	}
	producerArgs.Keys = []string{}
	producerArgs.Values = []string{}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
	defer h.close()
}

func (h *WebHandler) close() {
	//Wait for stats collector to finish
	time.Sleep(time.Second * 2)
	for _, prod := range h.Producers {
		prod.Close()
	}
}

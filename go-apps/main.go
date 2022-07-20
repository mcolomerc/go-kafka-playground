package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"mcolomerc/kafka-playground/config"
	"mcolomerc/kafka-playground/confluent"
	"mcolomerc/kafka-playground/franz"
	"mcolomerc/kafka-playground/producer"
	"mcolomerc/kafka-playground/sarama"
	"mcolomerc/kafka-playground/webserver"
	"net/http"

	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

//go:embed webapp/dist
var content embed.FS

func main() {
	// Load configuration
	config, err := config.LoadConfig("./.env")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println(config)

	if KafkaClient() {

		franz, err := franz.New(config, "franz")
		if err != nil {
			log.Fatal(err)
		}
		cflt, err := confluent.New(config, "confluent")
		if err != nil {
			log.Fatal(err)
		}

		sarama, err := sarama.New(config, "sarama")
		if err != nil {
			log.Fatal(err)
		}
		mux := http.NewServeMux()

		// Use the mux.Handle() function to register the file server as the handler for
		// all URL paths that start with "/static/". For matching paths, we strip the
		// "/static" prefix before the request reaches the file server.
		mux.Handle("/", staticHandler())

		producers := []producer.Producer{franz, cflt, sarama}

		webHandler := webserver.New(config, producers)

		mux.Handle("/api/produce", http.HandlerFunc(webHandler.Produce))

		mux.Handle("/metrics", promhttp.Handler())

		port := config.WEB_PORT
		log.Printf("Listening on: %d \n", port)
		if err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux); err != nil {
			log.Println(err)
			os.Exit(-1)
		}
	} else {
		log.Println("Kafka not running")
		os.Exit(-1)
	}
}

func KafkaClient() bool {
	var retries int = 5

	for retries > 0 {
		_, err := http.Get("http://broker:8090/kafka/v3/clusters")
		if err != nil {
			log.Println(err)
			time.Sleep(time.Second * 5)
			retries -= 1
		} else {
			break
		}
	}
	if retries == 0 {
		return false
	}
	return true
}

func staticHandler() http.Handler {
	fsys := fs.FS(content)
	html, _ := fs.Sub(fsys, "webapp/dist")

	return http.FileServer(http.FS(html))
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

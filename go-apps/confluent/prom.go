package confluent

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type StatsCollector struct {
	LastStats   Stats //This is the last stats object received from the channel
	Txmsgs      *prometheus.Desc
	Replyq      *prometheus.Desc
	Tx          *prometheus.Desc
	Txbytes     *prometheus.Desc
	Txmsg_bytes *prometheus.Desc
	Msg_cnt     *prometheus.Desc
	Msg_size    *prometheus.Desc
	Rx          *prometheus.Desc
	Rxbytes     *prometheus.Desc
}

func NewStatsCollector() *StatsCollector {
	return &StatsCollector{
		Txmsgs: prometheus.NewDesc("librdkafka_txmsgs",
			"Total number of messages transmitted (produced) to Kafka brokers",
			nil, nil,
		),
		Replyq: prometheus.NewDesc("librdkafka_replyq",
			"Number of ops (callbacks, events, etc) waiting in queue for application to serve with rd_kafka_poll()",
			nil, nil,
		),
		Tx: prometheus.NewDesc("librdkafka_tx", "Total number of requests sent to Kafka brokers",
			nil, nil,
		),
		Txbytes: prometheus.NewDesc("librdkafka_txbytes", "Total number of bytes transmitted to Kafka brokers",
			nil, nil,
		),
		Txmsg_bytes: prometheus.NewDesc("librdkafka_txmsg_bytes",
			"Total number of message bytes (including framing, such as per-Message framing and MessageSet/batch framing) transmitted to Kafka brokers",
			nil, nil,
		),
		Msg_cnt:  prometheus.NewDesc("librdkafka_msg_cnt", "Current number of messages in producer queues", nil, nil),
		Msg_size: prometheus.NewDesc("librdkafka_msg_size", "Current total size of messages in producer queues", nil, nil),
		Rx:       prometheus.NewDesc("librdkafka_rx", "Total number of responses received from Kafka brokers", nil, nil),
		Rxbytes:  prometheus.NewDesc("librdkafka_rxbytes", "Total number of bytes received from Kafka brokers", nil, nil),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *StatsCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.Txmsgs
	ch <- collector.Replyq
	ch <- collector.Tx
	ch <- collector.Txbytes
	ch <- collector.Txmsg_bytes
	ch <- collector.Msg_cnt
	ch <- collector.Msg_size
	ch <- collector.Rx
	ch <- collector.Rxbytes

}

//Collect implements required collect function for all promehteus collectors
func (collector *StatsCollector) Collect(ch chan<- prometheus.Metric) {

	select {
	case stats := <-GetMetrics():
		collector.LastStats = stats
		//GAUGEs
		m2 := prometheus.MustNewConstMetric(collector.Replyq, prometheus.GaugeValue, float64(stats.Replyq))
		m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
		ch <- m2
		m3 := prometheus.MustNewConstMetric(collector.Msg_cnt, prometheus.GaugeValue, float64(stats.MsgCnt))
		m3 = prometheus.NewMetricWithTimestamp(time.Now(), m3)
		ch <- m3
		m4 := prometheus.MustNewConstMetric(collector.Msg_size, prometheus.GaugeValue, float64(stats.MsgSize))
		m4 = prometheus.NewMetricWithTimestamp(time.Now(), m4)
		ch <- m4
		// Counters
		ch <- prometheus.MustNewConstMetric(collector.Txmsgs, prometheus.CounterValue, float64(stats.Txmsgs))
		ch <- prometheus.MustNewConstMetric(collector.Tx, prometheus.CounterValue, float64(stats.Tx))
		ch <- prometheus.MustNewConstMetric(collector.Txbytes, prometheus.CounterValue, float64(stats.TxBytes))
		ch <- prometheus.MustNewConstMetric(collector.Txmsg_bytes, prometheus.CounterValue, float64(stats.TxmsgBytes))
		ch <- prometheus.MustNewConstMetric(collector.Rx, prometheus.CounterValue, float64(stats.Rx))
		ch <- prometheus.MustNewConstMetric(collector.Rxbytes, prometheus.CounterValue, float64(stats.RxBytes))
	default:
		ch <- prometheus.MustNewConstMetric(collector.Txmsgs, prometheus.CounterValue, float64(collector.LastStats.Txmsgs))
		ch <- prometheus.MustNewConstMetric(collector.Tx, prometheus.CounterValue, float64(collector.LastStats.Tx))
		ch <- prometheus.MustNewConstMetric(collector.Txbytes, prometheus.CounterValue, float64(collector.LastStats.TxBytes))
		ch <- prometheus.MustNewConstMetric(collector.Txmsg_bytes, prometheus.CounterValue, float64(collector.LastStats.TxmsgBytes))
		ch <- prometheus.MustNewConstMetric(collector.Rx, prometheus.CounterValue, float64(collector.LastStats.Rx))
		ch <- prometheus.MustNewConstMetric(collector.Rxbytes, prometheus.CounterValue, float64(collector.LastStats.RxBytes))

		m2 := prometheus.MustNewConstMetric(collector.Replyq, prometheus.GaugeValue, float64(0))
		m2 = prometheus.NewMetricWithTimestamp(time.Now(), m2)
		ch <- m2
		m3 := prometheus.MustNewConstMetric(collector.Msg_cnt, prometheus.GaugeValue, float64(0))
		m3 = prometheus.NewMetricWithTimestamp(time.Now(), m3)
		ch <- m3
		m4 := prometheus.MustNewConstMetric(collector.Msg_size, prometheus.GaugeValue, float64(0))
		m4 = prometheus.NewMetricWithTimestamp(time.Now(), m4)
		ch <- m4
	}
}

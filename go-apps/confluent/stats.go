package confluent

type Stats struct {
	Name             string            `json:"name"`
	ClientID         string            `json:"client_id"`
	Type             string            `json:"type"`
	Ts               int64             `json:"ts"`
	Time             int               `json:"time"`
	Age              int               `json:"age"`
	Replyq           int               `json:"replyq"`
	MsgCnt           int               `json:"msg_cnt"`
	MsgSize          int               `json:"msg_size"`
	MsgMax           int               `json:"msg_max"`
	MsgSizeMax       int               `json:"msg_size_max"`
	SimpleCnt        int               `json:"simple_cnt"`
	MetadataCacheCnt int               `json:"metadata_cache_cnt"`
	Tx               int               `json:"tx"`
	TxBytes          int               `json:"tx_bytes"`
	Rx               int               `json:"rx"`
	RxBytes          int               `json:"rx_bytes"`
	Txmsgs           int               `json:"txmsgs"`
	TxmsgBytes       int               `json:"txmsg_bytes"`
	Rxmsgs           int               `json:"rxmsgs"`
	RxmsgBytes       int               `json:"rxmsg_bytes"`
	Brokers          map[string]Broker `json:"brokers"`
	Topics           map[string]Topic  `json:"topics"`
}

type Broker struct {
	Name           string `json:"name"`
	Nodeid         int    `json:"nodeid"`
	Nodename       string `json:"nodename"`
	Source         string `json:"source"`
	State          string `json:"state"`
	Stateage       int    `json:"stateage"`
	OutbufCnt      int    `json:"outbuf_cnt"`
	OutbufMsgCnt   int    `json:"outbuf_msg_cnt"`
	WaitrespCnt    int    `json:"waitresp_cnt"`
	WaitrespMsgCnt int    `json:"waitresp_msg_cnt"`
	Tx             int    `json:"tx"`
	Txbytes        int    `json:"txbytes"`
	Txerrs         int    `json:"txerrs"`
	Txretries      int    `json:"txretries"`
	Txidle         int    `json:"txidle"`
	ReqTimeouts    int    `json:"req_timeouts"`
	Rx             int    `json:"rx"`
	Rxbytes        int    `json:"rxbytes"`
	Rxerrs         int    `json:"rxerrs"`
	Rxcorriderrs   int    `json:"rxcorriderrs"`
	Rxpartial      int    `json:"rxpartial"`
	Rxidle         int    `json:"rxidle"`
	ZbufGrow       int    `json:"zbuf_grow"`
	BufGrow        int    `json:"buf_grow"`
	Wakeups        int    `json:"wakeups"`
	Connects       int    `json:"connects"`
	Disconnects    int    `json:"disconnects"`
	IntLatency     struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"int_latency"`
	OutbufLatency struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"outbuf_latency"`
	Rtt struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"rtt"`
	Throttle struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"throttle"`
	Req struct {
		Produce                             int `json:"Produce"`
		ListOffsets                         int `json:"ListOffsets"`
		Metadata                            int `json:"Metadata"`
		FindCoordinator                     int `json:"FindCoordinator"`
		SaslHandshake                       int `json:"SaslHandshake"`
		APIVersion                          int `json:"ApiVersion"`
		InitProducerID                      int `json:"InitProducerId"`
		AddPartitionsToTxn                  int `json:"AddPartitionsToTxn"`
		AddOffsetsToTxn                     int `json:"AddOffsetsToTxn"`
		EndTxn                              int `json:"EndTxn"`
		TxnOffsetCommit                     int `json:"TxnOffsetCommit"`
		SaslAuthenticate                    int `json:"SaslAuthenticate"`
		OffsetDeleteRequest                 int `json:"OffsetDeleteRequest"`
		DescribeClientQuotasRequest         int `json:"DescribeClientQuotasRequest"`
		AlterClientQuotasRequest            int `json:"AlterClientQuotasRequest"`
		DescribeUserScramCredentialsRequest int `json:"DescribeUserScramCredentialsRequest"`
	} `json:"req"`
}

type Topic struct {
	Topic       string `json:"topic"`
	Age         int    `json:"age"`
	MetadataAge int    `json:"metadata_age"`
	Batchsize   struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"batchsize"`
	Batchcnt struct {
		Min        int `json:"min"`
		Max        int `json:"max"`
		Avg        int `json:"avg"`
		Sum        int `json:"sum"`
		Stddev     int `json:"stddev"`
		P50        int `json:"p50"`
		P75        int `json:"p75"`
		P90        int `json:"p90"`
		P95        int `json:"p95"`
		P99        int `json:"p99"`
		P9999      int `json:"p99_99"`
		Outofrange int `json:"outofrange"`
		Hdrsize    int `json:"hdrsize"`
		Cnt        int `json:"cnt"`
	} `json:"batchcnt"`
	Partitions map[string]Partition `json:"partitions"`
}

type Partition struct {
	Partition         int    `json:"partition"`
	Broker            int    `json:"broker"`
	Leader            int    `json:"leader"`
	Desired           bool   `json:"desired"`
	Unknown           bool   `json:"unknown"`
	MsgqCnt           int    `json:"msgq_cnt"`
	MsgqBytes         int    `json:"msgq_bytes"`
	XmitMsgqCnt       int    `json:"xmit_msgq_cnt"`
	XmitMsgqBytes     int    `json:"xmit_msgq_bytes"`
	FetchqCnt         int    `json:"fetchq_cnt"`
	FetchqSize        int    `json:"fetchq_size"`
	FetchState        string `json:"fetch_state"`
	QueryOffset       int    `json:"query_offset"`
	NextOffset        int    `json:"next_offset"`
	AppOffset         int    `json:"app_offset"`
	StoredOffset      int    `json:"stored_offset"`
	CommitedOffset    int    `json:"commited_offset"`
	CommittedOffset   int    `json:"committed_offset"`
	EOFOffset         int    `json:"eof_offset"`
	LoOffset          int    `json:"lo_offset"`
	HiOffset          int    `json:"hi_offset"`
	LsOffset          int    `json:"ls_offset"`
	ConsumerLag       int    `json:"consumer_lag"`
	ConsumerLagStored int    `json:"consumer_lag_stored"`
	Txmsgs            int    `json:"txmsgs"`
	Txbytes           int    `json:"txbytes"`
	Rxmsgs            int    `json:"rxmsgs"`
	Rxbytes           int    `json:"rxbytes"`
	Msgs              int    `json:"msgs"`
	RxVerDrops        int    `json:"rx_ver_drops"`
	MsgsInflight      int    `json:"msgs_inflight"`
	NextAckSeq        int    `json:"next_ack_seq"`
	NextErrSeq        int    `json:"next_err_seq"`
	AckedMsgid        int    `json:"acked_msgid"`
}

var metrics Metrics

func init() {
	metrics = Metrics{StatsChan: make(chan Stats)}
}

type Metrics struct {
	StatsChan chan Stats
}

func GetMetrics() chan Stats {
	return metrics.StatsChan
}

func ProcessStats(stats Stats) {
	metrics.StatsChan <- stats // send to channel
}

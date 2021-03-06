package nsqd

import (
	"crypto/md5"
	"hash/crc32"
	"io"
	"log"
	"os"
	"time"
)

type nsqdOptions struct {
	// basic options
	ID                     int64    `flag:"worker-id" cfg:"id"`
	Verbose                bool     `flag:"verbose"`
	TCPAddress             string   `flag:"tcp-address"`
	HTTPAddress            string   `flag:"http-address"`
	HTTPSAddress           string   `flag:"https-address"`
	BroadcastAddress       string   `flag:"broadcast-address"`
	NSQLookupdTCPAddresses []string `flag:"lookupd-tcp-address" cfg:"nsqlookupd_tcp_addresses"`
	AuthHTTPAddresses      []string `flag:"auth-http-address" cfg:"auth_http_addresses"`

	// diskqueue options
	DataPath        string        `flag:"data-path"`
	MemQueueSize    int64         `flag:"mem-queue-size"`
	MaxBytesPerFile int64         `flag:"max-bytes-per-file"`
	SyncEvery       int64         `flag:"sync-every"`
	SyncTimeout     time.Duration `flag:"sync-timeout"`

	// msg and command options
	MsgTimeout    time.Duration `flag:"msg-timeout" arg:"1ms"`
	MaxMsgTimeout time.Duration `flag:"max-msg-timeout"`
	MaxMsgSize    int64         `flag:"max-msg-size" deprecated:"max-message-size" cfg:"max_msg_size"`
	MaxBodySize   int64         `flag:"max-body-size"`
	MaxReqTimeout time.Duration `flag:"max-req-timeout"`
	ClientTimeout time.Duration

	// client overridable configuration options
	MaxHeartbeatInterval   time.Duration `flag:"max-heartbeat-interval"`
	MaxRdyCount            int64         `flag:"max-rdy-count"`
	MaxOutputBufferSize    int64         `flag:"max-output-buffer-size"`
	MaxOutputBufferTimeout time.Duration `flag:"max-output-buffer-timeout"`

	// statsd integration
	StatsdAddress  string        `flag:"statsd-address"`
	StatsdPrefix   string        `flag:"statsd-prefix"`
	StatsdInterval time.Duration `flag:"statsd-interval" arg:"1s"`
	StatsdMemStats bool          `flag:"statsd-mem-stats"`

	// e2e message latency
	E2EProcessingLatencyWindowTime  time.Duration `flag:"e2e-processing-latency-window-time"`
	E2EProcessingLatencyPercentiles []float64     `flag:"e2e-processing-latency-percentile" cfg:"e2e_processing_latency_percentiles"`

	// TLS config
	TLSCert             string `flag:"tls-cert"`
	TLSKey              string `flag:"tls-key"`
	TLSClientAuthPolicy string `flag:"tls-client-auth-policy"`
	TLSRootCAFile       string `flag:"tls-root-ca-file"`
	TLSRequired         bool   `flag:"tls-required"`

	// compression
	DeflateEnabled  bool `flag:"deflate"`
	MaxDeflateLevel int  `flag:"max-deflate-level"`
	SnappyEnabled   bool `flag:"snappy"`
}

func NewNSQDOptions() *nsqdOptions {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	o := &nsqdOptions{
		TCPAddress:       "0.0.0.0:4150",
		HTTPAddress:      "0.0.0.0:4151",
		HTTPSAddress:     "0.0.0.0:4152",
		BroadcastAddress: hostname,

		MemQueueSize:    10000,
		MaxBytesPerFile: 104857600,
		SyncEvery:       2500,
		SyncTimeout:     2 * time.Second,

		MsgTimeout:    60 * time.Second,
		MaxMsgTimeout: 15 * time.Minute,
		MaxMsgSize:    1024768,
		MaxBodySize:   5 * 1024768,
		MaxReqTimeout: 1 * time.Hour,
		ClientTimeout: 60 * time.Second,

		MaxHeartbeatInterval:   60 * time.Second,
		MaxRdyCount:            2500,
		MaxOutputBufferSize:    64 * 1024,
		MaxOutputBufferTimeout: 1 * time.Second,

		StatsdPrefix:   "nsq.%s",
		StatsdInterval: 60 * time.Second,
		StatsdMemStats: true,

		E2EProcessingLatencyWindowTime: time.Duration(10 * time.Minute),

		DeflateEnabled:  true,
		MaxDeflateLevel: 6,
		SnappyEnabled:   true,
	}

	h := md5.New()
	io.WriteString(h, hostname)
	o.ID = int64(crc32.ChecksumIEEE(h.Sum(nil)) % 1024)

	return o
}

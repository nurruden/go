package common

import "time"

type CollectConfig struct {
	Path string `json:"path"`
	ModuleName string `json:"module_name"`
	Topic string `json:"topic"`
}

type CollectSystemInfoConfig struct {
	Interval time.Duration 	`json:"interval"`
	Topic string    `json:"topic"`
}

type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	LogConf LogConfig `ini:"logs"`
	EtcdConf EtcdConfig `ini:"etcd"`
}

type KafkaConfig struct {
	Address string `ini:"address"`
	QueueSize int  `ini:"queue_size"`
}

type EtcdConfig struct {
	Address string `ini:"address"`
	EtcdKey string `ini:"etcd_key"`
	EtcdCollectSystemInfoKey string `ini:"etcd_collect_system_info_key"`
}

type LogConfig struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType string `ini:"log_type"`
	Module string `ini:"module"`
}


type LogAgentData struct {
	IP string
	Data string
}

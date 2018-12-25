package main

import (
	"allanlogagent/kafka"
	"allanlogagent/tailf"
	"fmt"
	"oconfig"
	"strings"
	"xlog"
)

var (
	appConfig AppConfig
)

type AppConfig struct {
	KafkaConf KafkaConfig `ini:"kafka"`
	CollectLogsConf CollectLogsConfig	`ini:"collect_log_conf""`
	LogConf LogConfig `ini:"logs"`
}

type KafkaConfig struct {
	Address	string `ini:"address"`
	QueueSize int `ini:queue_size`
}

type CollectLogsConfig struct{
	LogFilenames string	`ini:"log_filenames"`

}

type LogConfig struct {
	LogLevel string `ini:"log_level"`
	Filename string `ini:"filename"`
	LogType string `ini:"log_type"`
	Module string `ini:"module"`

}

func initConfig(filename string) (err error) {
	err = oconfig.UnMarshalFile(filename,&appConfig)
	if err != nil {
		return
	}

	xlog.LogDebug("Read configuration successfully, configuration:%#v\n",appConfig)

	return
}

func run()(err error) {
	//read log data from tailf and send it out with kafka

	for {
		//read log data from tailf
		line, err := tailf.ReadLine()
		if err != nil {
			continue
		}
		xlog.LogDebug("Line:%s",line.Text)
		msg := &kafka.Message{
			Line: line.Text,
			Topic: "nginx_log",
		}
		err = kafka.SendLog(msg)
		if err != nil {
			xlog.LogWarn("Kafka send log failed, err:%v",err)
		}
		xlog.LogDebug("Send to kafka successfully")
	}
	return
}

func initLog() (err error){
	var logType int
	var level int
	if appConfig.LogConf.LogType == "console" {
		logType = xlog.XLogTypeConsole
	}else {
		logType = xlog.XLogTypeFile
	}

	switch appConfig.LogConf.LogLevel {
	case "debug":
		level = xlog.XLogLevelDebug
	case "trace":
		level = xlog.XLogLevelTrace
	case "info":
		level = xlog.XLogLevelInfo
	case "warn":
		level = xlog.XLogLevelWarn
	case "error":
		level = xlog.XLogLevelError
	default:
		level = xlog.XLogLevelDebug
	}

	err = xlog.Init(logType,level,appConfig.LogConf.Filename,appConfig.LogConf.Module)
	return

}
func main() {
	err := initConfig("./conf/config.ini")
	if err != nil {
		panic(fmt.Sprintf("init config failed,err:%v",err))
	}

	err = initLog()
	if err != nil {
		panic(fmt.Sprintf("init logs failed, err:%v",err))
	}
	xlog.LogDebug("Init log successfully!")


	address := strings.Split(appConfig.KafkaConf.Address,",")
	err = kafka.Init(address,appConfig.KafkaConf.QueueSize)
	if err != nil {
		panic(fmt.Sprintf("Init kafka client failed,err:%v",err))
	}
	xlog.LogDebug("Init kafka successfully")

	err = tailf.Init(appConfig.CollectLogsConf.LogFilenames)
	if err != nil {
		panic(fmt.Sprintf("Init tailf client failed,err:%v",err))
	}
	xlog.LogDebug("Init tailf successfully")

	err = run()
	if err != nil {
		xlog.LogError("Run failed,err: %v",err)
	}
	xlog.LogDebug("run finished")
}


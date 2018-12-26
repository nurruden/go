package main

import (
	"allanlogagent/common"
	"allanlogagent/etcd"
	"allanlogagent/kafka"
	"allanlogagent/tailf"
	"fmt"
	"oconfig"
	"strings"
	"xlog"
)

var (
	appConfig common.AppConfig
)



func initConfig(filename string) (err error) {
	err = oconfig.UnMarshalFile(filename,&appConfig)
	if err != nil {
		return
	}

	xlog.LogDebug("Read configuration successfully, configuration:%#v",appConfig)

	return
}

func run()(err error) {
	//read log data from tailf and send it out with kafka
	//Inspect etcd continually, if yes, then manage log collect
	tailf.Run()
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

    //Init etcd client
	address = strings.Split(appConfig.EtcdConf.Address,",")
	err = etcd.Init(address,appConfig.EtcdConf.EtcdKey)
	if err != nil {
		panic(fmt.Sprintf("Init etcd client failed,err:%v",err))
	}
	xlog.LogDebug("Init etcd successfully address:%v",address)

	logCollectConf,err := etcd.GetConfig(appConfig.EtcdConf.EtcdKey)
	xlog.LogDebug("etcd conf:%#v",logCollectConf)
    watch := etcd.Watch()
	err = tailf.Init(logCollectConf,watch)
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


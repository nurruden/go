package main

import (
	"allanlogagent/common"
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)


func main()  {

	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
		DialTimeout:3 * time.Second,
	})

	defer client.Close()
    localIP,err := common.GetLocalIP()
    if err != nil {
    	fmt.Printf("Get local ip failed,err:%v",err)
	}

    fmt.Printf("Get local ip succ,ip:%v",localIP)
	var logCollectArray []common.CollectConfig
	logCollect := common.CollectConfig{
		Topic:"nginx_log",
		Path:"/Users/northstar/Documents/go/src/allanlogagent/logs/my.log",
		ModuleName:"nginx",
	}
	logCollectArray = append(logCollectArray, logCollect)


	//logCollectArray = append(logCollectArray, logCollect)
	//
	//logCollect = common.CollectConfig{
	//	Topic:"nginx_log",
	//	Path:"c:/tmp/c.log",
	//	ModuleName:"nginx",
	//}
	//
	//logCollectArray = append(logCollectArray, logCollect)

	data, err := json.Marshal(logCollectArray)
	if err != nil {
		fmt.Printf("marshal failed, conf:%#v\n", logCollectArray)
		return
	}

	etcdKey := fmt.Sprintf("/allanlogagent/%s/conf",localIP)
	_, err = client.Put(context.Background(), etcdKey, string(data))
	if err != nil {
		fmt.Printf("put failed, err:%v\n", err)
	}
}
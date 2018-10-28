package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("Connect failed, err:", err)
		return
	}
	fmt.Println("Conect succeed")
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "/logagent/conf", "sample_value")
	cancel()
	if err != nil {
		fmt.Println("Put failed, err:", err)
		return

	}
	ctx,cancel = context.WithTimeout(context.Background(),time.Second)
	resp, err := cli.Get(ctx,"/logagent/conf/")
	cancel()
	if err != nil {
		fmt.Println("Get failed, err:", err)
		return
	}
	for _,ev := range resp.Kvs{
		fmt.Printf("%s: %s\n",ev.Key,ev.Value)
	}
}
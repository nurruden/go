package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
	"time"
)

func initRedis()(conn redis.Conn,err error) {
	//没有使用连接池
	conn,err = redis.Dial("tcp","localhost:6379")
	if err != nil {
		fmt.Printf("conn redis failed,err:%v\n",err)
		return
	}
	return
}

func testSetGet(conn redis.Conn){
	key := "abc"
	_,err := conn.Do("set",key,"this is a test")
	if err != nil{
		fmt.Printf("set failed,err:%v\n",err)
		return
	}
	//reply,err := conn.Do("get",key)

	data,err := redis.String(conn.Do("get",key))
	if err != nil {
		fmt.Printf("get failed,err:%v\n",err)
		return
	}
	fmt.Printf("key:%s value:%s\n",key,data)
}

func testHSetGet(conn redis.Conn){
	key := "abc"
	_,err := conn.Do("hset","books",key,"this is a test")
	if err != nil{
		fmt.Printf("set failed,err:%v\n",err)
		return
	}
	//reply,err := conn.Do("get",key)

	data,err := redis.String(conn.Do("hget","books",key))
	if err != nil {
		fmt.Printf("get failed,err:%v\n",err)
		return
	}
	fmt.Printf("key:%s value:%s\n",key,data)
}

func testMSetGet(conn redis.Conn){
	key := "abc"
	key1 := "efg"
	_,err := conn.Do("mset",key,"this is a test",key1,"dkakdslad")
	if err != nil{
		fmt.Printf("set failed,err:%v\n",err)
		return
	}
	//reply,err := conn.Do("get",key)

	data,err := redis.Strings(conn.Do("mget",key,key1))
	if err != nil {
		fmt.Printf("get failed,err:%v\n",err)
		return
	}
	for _,val := range data{
	    fmt.Printf("value:%s\n",val)
	}
}

func testExpire(conn redis.Conn){
	key := "abc"
	_,err := conn.Do("set",key,"this is a test")
	if err != nil{
		fmt.Printf("set failed,err:%v\n",err)
		return
	}
	//reply,err := conn.Do("get",key)
    //defer conn.Close()
    _,err = conn.Do("expire","abc",10)
	//data,err := redis.String(conn.Do("get",key))
	if err != nil {
		fmt.Printf("get failed,err:%v\n",err)
		return
	}
	//fmt.Printf("key:%s value:%s\n",key,data)
}

func newPool(serverAddr string,passwd string)(pool *redis.Pool){
	return &redis.Pool{

		MaxIdle:16,
		MaxActive:1024,
		IdleTimeout:240*time.Second,
		Dial: func()(redis.Conn,error){
			conn,err :=redis.Dial("tcp",serverAddr)
			redis.Dial("tcp",serverAddr)
			if err!= nil {
				return nil,err
			}

			if len(passwd)>0{
				_,err = conn.Do("auth",passwd)
				if err != nil {
					return nil,err
				}
			}
			return conn,err
		},

		TestOnBorrow: func(c redis.Conn,t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_,err := c.Do("ping")
			return err
		},
	}
}

func testRedisPool(){
	pool := newPool("localhost:6379","")

	conn := pool.Get()
	conn.Do("set","abc","233444445676543")

	val,err := redis.String(conn.Do("get","abc"))
	fmt.Printf("val: %s err:%v\n",val,err)
	//把连接归还到连接池
	conn.Close()

}

func main() {
	conn,err := initRedis()

	if err != nil {
		return
	}
    _=conn
	//testSetGet(conn)
	//testHSetGet(conn)
	//testMSetGet(conn)
	//testExpire(conn)
	testRedisPool()
}


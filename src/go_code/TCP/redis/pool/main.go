package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//初始化连接

func init() {
	pool = &redis.Pool{
		MaxActive:   0,
		MaxIdle:     8,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {
	//先从pool里面取出一个链接
	conn := pool.Get()
	defer conn.Close()
	_, err := conn.Do("set", "name", "tom")
	if err != nil {
		fmt.Println("bad err:", err)
	}
	//取出
	name, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("bad")
	}
	fmt.Println(name)
	pool.Close()
}

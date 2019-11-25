package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

//初始化
func init() {
	pool = &redis.Pool{
		MaxActive:   8,
		MaxIdle:     0,
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

func main() {

	//先从pool 池中取出一个连接
	//get方法与指针绑定
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("set", "actor", "tom")
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	r, err := redis.String(c.Do("get", "actor"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(r)
	defer pool.Close() //一旦关闭就无法开启
}

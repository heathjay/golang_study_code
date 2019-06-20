package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect err:", err)
		return
	}
	//fmt.Println("connect good", conn)

	//向redis里面写入set
	_, err = conn.Do("set", "id", "kkkkk")
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}
	fmt.Println("do good")

	res, err := redis.String(conn.Do("get", "id"))
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}
	fmt.Println("res:", res)

	_, err = conn.Do("hset", "user1", "name", "tom")
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}

	_, err = conn.Do("hset", "user1", "age", 11)
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}

	user1age, err := redis.String(conn.Do("hget", "user1", "age"))
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}
	fmt.Println("user1age:", user1age)

	//res 是一个接口
	//redis提供了自身的方法可以进行转换

	//批量set和get
	_, err = conn.Do("hmset", "user2", "name", "joo", "age", 1)
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}

	hmget, err := redis.Strings(conn.Do("hmget", "user2", "name", "age"))
	if err != nil {
		fmt.Println("do connect:", err)
		return
	}
	for v, i := range hmget {
		fmt.Println("get :", i, ":", v)
	}

}

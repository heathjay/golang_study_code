package main

//1.安装第三方redis库
//2.提供redis的api
import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//do 返回一个接口

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect error")
		return
	}
	//_, err = c.Do("set", "name1", "tommoy")
	_, err = c.Do("hmset", "user01", "name", "tommy")
	if err != nil {
		fmt.Println("set error")
		return
	}
	_, err = c.Do("hmset", "user01", "age", "18")
	if err != nil {
		fmt.Println("set error")
		return
	}
	//1.单个字符串
	//2.多个字符串 redis.Strings
	//3.hash, do("hset", "user01", "name", "tom")

	r, err := redis.String(c.Do("hget", "user01", "age"))
	if err != nil {
		fmt.Println("get error")
		return
	}
	fmt.Println(r)
	r, err = redis.String(c.Do("hget", "user01", "name"))
	if err != nil {
		fmt.Println("get error")
		return
	}
	fmt.Println(r)
	r1, err := redis.Strings(c.Do("hmget", "user01", "name", "age"))
	if err != nil {
		fmt.Println("get error")
		return
	}
	fmt.Printf("r=%v\n", r1) //循环输出
	for i, v := range r1 {
		fmt.Printf("r[%d] = %s\n", i, v)
	}
	defer c.Close()
}

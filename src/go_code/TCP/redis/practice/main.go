package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//monster name,skill,age

	fmt.Println("monster compose")

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("conn bad: ", err)
		return
	}
	fmt.Println("conn succ: ")
	fmt.Println("monster name:")
	var name, age, skill string
	fmt.Scanln(&name)
	fmt.Println("age:")
	fmt.Scanln(&age)
	fmt.Println("skill")
	fmt.Scanln(&skill)

	_, err = conn.Do("hmset", "user1", "name", name, "age", age, "skill", skill)

	if err != nil {
		fmt.Println("input bad:", err)
	}
	fmt.Println("input suc:")
	res, err := redis.Strings(conn.Do("hmget", "user1", "name", "age", "skill"))
	if err != nil {
		fmt.Println("bad")
		return
	}
	for i, v := range res {
		fmt.Printf("s[%v]: %v\n", i, v)
	}
}

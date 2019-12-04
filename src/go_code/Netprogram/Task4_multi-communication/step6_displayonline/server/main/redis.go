package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

/*5.2进行redis初始化工作
 */
var pool *redis.Pool

//这个函数不会自动调用
func initPool(address string, maxId int, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle:     maxId,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}

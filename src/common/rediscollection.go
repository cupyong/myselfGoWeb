package common

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

var c, err = redis.Dial("tcp", "127.0.0.1:6379")


//设置缓存
func SetCache(key,value string)  {
	_, err = c.Do("SET", key, value, "EX", "60")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
}

//获取缓存
func GetCache(key string) (string,error) {
	value, err := redis.String(c.Do("GET", key))
	return value,err
}





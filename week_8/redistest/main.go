package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"strings"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	_ = initClient()
	insertData(10)
	insertData(20)
	insertData(50)
	insertData(100)
	insertData(200)
	insertData(1024)
	insertData(1024 * 5)
}

func insertData(size int) {
	_, err := rdb.FlushAll().Result()
	if err != nil {
		fmt.Printf("flush redis error %s", err.Error())
		return
	}
	before := getRedisInfo()
	data := make([]byte, size)
	for i := 0; i < 10000; i++ {
		_, err = rdb.Set(fmt.Sprintf("%4d", i), data, 0).Result()
		if err != nil {
			fmt.Printf("insert error %s\n", err.Error())
			break
		}
	}
	after := getRedisInfo()

	fmt.Printf("value size %d bytes, before usage: %d bytes, after usage %d bytes , perusage: %f bytes\n", size, before, after, float64(after-before)/10000)
}

func getRedisInfo() int64 {
	info, err := rdb.Info("Memory").Result()
	if err != nil {
		fmt.Printf("get redis info error %s\n", err.Error())
		return 0
	} else {
		memo := strings.Replace(strings.Split(info, "\r\n")[1], "used_memory:", "", 1)
		m, _ := strconv.ParseInt(memo, 10, 64)
		return m
	}
}

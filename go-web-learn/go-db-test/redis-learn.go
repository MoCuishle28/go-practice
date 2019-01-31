package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/gomodule/redigo/redis"
)

var (
	Pool *redis.Pool
)


func init() {
	redisHost := ":6379"
	Pool = newPool(redisHost)
	close()
}


func newPool(server string) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		}}
}


func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}


func Get(key string) ([]byte, error) {
	conn := Pool.Get()	// 从连接池拿连接
	defer conn.Close()

	var data []byte 	// data是 字节切片
	data, err := redis.Bytes(conn.Do("GET", key))	// 执行 API 与原声 redis 几乎一样(返回的是字节切片)
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}


func Set(key string, val string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("SET", key, val))
	if err != nil {
		return data, fmt.Errorf("error set key %s: %v", key, err)
	}
	return data, err
}


func main() {
	test, err := Get("test")
	fmt.Println("key = test 的字节切片是:", test, err)
	fmt.Println("key = test is", string(test))

	test, err = Set("test", "999999999")
	test, err = Get("test")
	fmt.Println("key = test is", string(test))	
}
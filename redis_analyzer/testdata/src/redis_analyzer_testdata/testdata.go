package redis_analyzer_testdata

import (
	"fmt"
	"github.com/go-redis/redis"
)

func NewConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务地址
		Password: "",               // Redis 密码
		DB:       0,                // Redis 数据库索引
	})

	return client
}

// Example1 示例1: 正常创建链接 + 正常关闭链接
// 预期: success
func Example1() {
	// 创建 Redis 客户端
	var client = NewConnection()

	// 设置 Redis 键值对
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取 Redis 键值对
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// 删除 Redis 键值对
	err = client.Del("key").Err()
	if err != nil {
		panic(err)
	}

	// 关闭 Redis 客户端
	err = client.Close()
	if err != nil {
		panic(err)
	}
}

// Example2 示例2: 正常创建链接 + 正常关闭链接
// 预期: failed
func Example2() {
	// 创建 Redis 客户端
	var client = NewConnection()

	// 设置 Redis 键值对
	err := client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 获取 Redis 键值对
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// 删除 Redis 键值对
	err = client.Del("key").Err()
	if err != nil {
		panic(err)
	}
}

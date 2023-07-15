package main

import (
	"context"
	redisRepo "github.com/punkestu/hello-redis/repo/redis"
	"log"
	"time"
)

func main() {
	client := redisRepo.NewConnection()

	err := client.SetValue(context.Background(), "testKey", "testValue", time.Duration(1)*time.Minute) // expired in time.Duration so we use time.Duration(duration) * scale
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("set operation success")

	res, err := client.GetValue(context.Background(), "testKey")
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("get operation success. result:", res)
}

package main

import (
	"myServer/cache"
	"myServer/config"
	"myServer/kafka"
	"myServer/log"
	"myServer/routes"
)

func main() {
	loading()
	// kafka.KafkaTest()
	cache.RedisTest()
	r := routes.NewRouter()
	err := r.Run(config.Config.System.HttpPort)
	if err != nil {
		panic(err)
	}
}

func loading() {
	config.InitConfig()
	kafka.InitKafka()
	cache.InitCache()
	err := log.InitLogger()
	if err != nil {
		panic(err)
	}
}

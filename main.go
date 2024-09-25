package main

import (
	"myServer/config"
	"myServer/kafka"
	"myServer/log"
	"myServer/routes"
)

func main() {
	loading()
	r := routes.NewRouter()
	err := r.Run(config.Config.System.HttpPort)
	if err != nil {
		panic(err)
	}
}

func loading() {
	config.InitConfig()
	kafka.InitKafka()
	err := log.InitLogger()
	if err != nil {
		panic(err)
	}
}

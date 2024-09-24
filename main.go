package main

import (
	"exercise/config"
	"exercise/kafka"
	"exercise/routes"
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
}

package config

import (
	"os"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	System *System               `yaml:"system"`
	Kafka  map[string]*KafkaConf `yaml:"Kafka"`
	Redis  *RedisConf            `yaml:"redis"`
	// Mysql
	// Redis
}

type System struct {
	AppEnv      string `yaml:"env"`
	Domain      string `yaml:"domain"`
	Version     string `yaml:"version"`
	HttpPort    string `yaml:"httpPort"`
	Host        string `yaml:"Host"`
	UploadModel string `yaml:"UploadModel"`
}

type KafkaConf struct {
	DisableConsumer bool   `yaml:"disableConsumer"`
	Debug           bool   `yaml:"debug"`
	Address         string `yaml:"address"`
	RequiredAck     int    `yaml:"requiredAck"`
	ReadTimeout     int64  `yaml:"readTimeout"`
	WriteTimeout    int64  `yaml:"writeTimeout"`
	MaxOpenRequests int    `yaml:"maxOpenRequests"`
	Partition       int    `yaml:"partition"`
}

type RedisConf struct {
	RedisHost     string `yaml:"redisHost"`
	RedisPort     string `yaml:"redisPort"`
	RedisUsername string `yaml:"redisUsername"`
	RedisPassword string `yaml:"redisPwd"`
	RedisDbName   int    `yaml:"redisDbName"`
	RedisNetwork  string `yaml:"redisNetwork"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/config/locales")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}
}

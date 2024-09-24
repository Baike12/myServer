package kafka

import "github.com/Shopify/sarama"

func middlewareConsumerHandler(fn func(message *sarama.ConsumerMessage) error) func(mesage *sarama.ConsumerMessage) error {
	return func(msg *sarama.ConsumerMessage) error {
		return fn(msg)
	}
}

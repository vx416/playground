package consumer

import (
	"github.com/Shopify/sarama"
)

func NewClient(addr []string, config *sarama.Config) (sarama.Client, error) {
	return sarama.NewClient(addr, config)
}

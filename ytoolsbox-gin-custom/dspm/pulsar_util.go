package dspm

import (
	"context"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"log"
	"main/cons"
	"time"
)

func CreateClient() pulsar.Client {
	options := pulsar.ClientOptions{
		URL: fmt.Sprintf("pulsar://%s:%v", cons.DspmAddr, cons.PulsarPort), // 替换为你的 Pulsar 服务地址
	}

	options.Authentication = pulsar.NewAuthenticationToken(cons.PulsarToken)
	client, err := pulsar.NewClient(options)
	if err != nil {
		log.Fatalf("connection pulsar fail err:%v", err)
	}
	return client
}

func CreateProducer(client pulsar.Client, topic string) pulsar.Producer {
	producer, err := client.CreateProducer(pulsar.ProducerOptions{
		Topic:                   topic,
		BatchingMaxPublishDelay: time.Duration(1000) * time.Millisecond,
		BatchingMaxMessages:     2000,
		MaxPendingMessages:      3000,
		SendTimeout:             time.Duration(10) * time.Second,
		CompressionType:         pulsar.LZ4,
		CompressionLevel:        pulsar.Default,
	})
	if err != nil {
		log.Fatalf("create producer fail err:%v", err)
	}
	return producer
}

func SendDataBatchToPulsar(producer pulsar.Producer, messages []*pulsar.ProducerMessage) {
	for _, message := range messages {
		producer.SendAsync(context.Background(), message, func(id pulsar.MessageID, msg *pulsar.ProducerMessage, err error) {
			if err != nil {
				log.Fatalf("Failed to publish message to Pulsar: %v", err)
			}
		})
	}
	log.Printf("send msg :%v success\n", len(messages))
}

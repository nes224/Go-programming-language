package main

import (
	"log"
	"time"

	"github.com/kafka/golang/config"
	"github.com/kafka/golang/models"
	"github.com/kafka/golang/pkg/utils"
	"github.com/segmentio/kafka-go"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn := utils.KafkConn(config)

	if !utils.IsTopicAlreadyExists(conn, config.Topic) {
		topicConfigs := []kafka.TopicConfig{
			{
				Topic:             config.Topic,
				NumPartitions:     1,
				ReplicationFactor: 1,
			},
		}

		err := conn.CreateTopics(topicConfigs...)
		if err != nil {
			panic(err.Error())
		}
	}

	// Mock data
	data := func() []kafka.Message {
		products := []models.Product{
			{
				Id:    "1111-1111-xiii-1111",
				Title: "Coffee",
			},
			{
				Id:    "2222-2222-iixi-2222",
				Title: "Tea",
			},
			{
				Id:    "3333-3333-viix-3333",
				Title: "Milk",
			},
		}

		messages := make([]kafka.Message, 0)
		for _, p := range products {
			messages = append(messages, kafka.Message{
				Value: utils.CompressToJsonBytes(&p),
			})
		}
		return messages
	}()

	conn.SetDeadline(time.Now().Add(10 * time.Second))
	conn.WriteMessages(data...)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

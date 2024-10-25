package main

import (
	"fmt"
	"log"

	"github.com/kafka/golang/config"
	"github.com/kafka/golang/pkg/utils"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config.")
	}
	conn := utils.KafkConn(config)

	for {
		message, err := conn.ReadMessage(10e3)
		if err != nil {
			break
		}
		fmt.Println(string(message.Value))
	}
	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}
}
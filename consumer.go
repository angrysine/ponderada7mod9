package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.mongodb.org/mongo-driver/mongo"
)

var ConsumerPointer *kafka.Consumer
var CollectionPointer *mongo.Collection

func Consumer(collectionPointer *mongo.Collection) *kafka.Consumer {
	
	CollectionPointer = collectionPointer 
	if ConsumerPointer == nil {
		GenerateConsumer()
		return ConsumerPointer
	}
	return ConsumerPointer
}

func GenerateConsumer() {
	// Configurações do consumidor
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092,localhost:39092",
		"group.id":          "go-consumer-group",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	// defer consumer.Close()

	ConsumerPointer = consumer
}

func Subscribe(consumer *kafka.Consumer,topic string) {
	fmt.Printf("aqui: %v",consumer)
	// Assinar tópico
	err := consumer.SubscribeTopics([]string{topic}, nil)

	throw(err)

	// Consumir mensagens
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message: %s\n", string(msg.Value))
			Writer("./logs/teste.txt",fmt.Sprintf("%v",string(msg.Value)))
			result := strings.Split(string(msg.Value), ",")
			age, _ := strconv.Atoi(result[2])
			hours_spent_value, _ := strconv.Atoi(result[3])
			data = &Data{name: result[0], password: result[1], age: age, hours_spent: hours_spent_value}
			fmt.Println("aqui5")
			Insert(db, *data)
			Writer("./logs/consumer_logs.txt", "name: "+data.name+" password: "+data.password+" age: "+strconv.Itoa(data.age)+" hours_spent: "+strconv.Itoa(data.hours_spent)+"\n")
			fmt.Println("aqui6")
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			break
		}
	}
}

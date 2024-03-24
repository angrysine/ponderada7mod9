package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	godotenv "github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	err := godotenv.Load(".env")
	throw(err)
	var password = os.Getenv("mongopassword")
	fmt.Println("aqui")
	client := Mongo(password)
	collection := GetCollection("betinhoDb", "users", client)
	producer := Producer()
	consumer := Consumer(collection)
	fmt.Println("aqui2")
	go RunProducer(producer)
	go Subscribe(consumer,"test_topic")
	select {}
}

func throw(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateMessage() string {
	name := "user" + strconv.Itoa(rand.IntN(100))
	password := "password" + strconv.Itoa(rand.IntN(100))
	age := rand.IntN(40)
	hours_spent := rand.IntN(100)
	text := name + "," + password + "," + strconv.Itoa(age) + "," + strconv.Itoa(hours_spent)
	return text
}

func RunProducer(producer *kafka.Producer) {

	for {
		Publish(GenerateMessage(),"test_topic",producer)
		time.Sleep(1000000000)
	}

}
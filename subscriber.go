package main

import (
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	name        string
	password    string
	age         int
	hours_spent int
}

var data *Data

var db *mongo.Collection

var messageToCompare []byte

var messagePubHandlerSub mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	var text = fmt.Sprintf("Recebido: %s do tópico: %s com QoS: %d\n", msg.Payload(), msg.Topic(), msg.Qos())
	fmt.Printf(text)
	var textBytes = msg.Payload()
	messageToCompare = textBytes
	
}

var connectHandlerSub mqtt.OnConnectHandler = func(client mqtt.Client) {
	Writer("subscriber_logs.txt", "connected"+"\n")
}

var connectLostHandlerSub mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	var text = fmt.Sprintf("Connection lost: %v", err)
	Writer("subscriber_logs.txt", text+"\n")
}

func Subscriber(dbPointer *mongo.Collection) {
	fmt.Println("Subscriber")
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}
	db = dbPointer

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	var connectionString = fmt.Sprintf("tls://%s%d/mqtt", broker, port)
	fmt.Println(connectionString)
	opts.AddBroker(connectionString)
	opts.SetClientID("Subscriber")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.SetDefaultPublishHandler(messagePubHandlerSub)
	opts.OnConnect = connectHandlerSub
	opts.OnConnectionLost = connectLostHandlerSub

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := client.Subscribe("test/topic", 1, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return
	}
}

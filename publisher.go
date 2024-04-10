package main

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	godotenv "github.com/joho/godotenv"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection lost: %v", err)
}

func Publisher() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s%d/mqtt", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		name := "user" + strconv.Itoa(rand.IntN(100))
		password := "password" + strconv.Itoa(rand.IntN(100))
		age := rand.IntN(40)
		hours_spent := rand.IntN(100)
		
		datajson,_ :=  json.Marshal(map[string]interface{}{ "name": name, "password": password, "age": age, "hours_spent": hours_spent})
		token := client.Publish("test_topic/fazol", 1, false, datajson)
		token.Wait()
		
		time.Sleep(2 * time.Second)
	}
}

func PublisherOne() Data{
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %s", err)
	}

	var broker = os.Getenv("BROKER_ADDR")
	var port = 8883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tls://%s%d/mqtt", broker, port))
	opts.SetClientID("Publisher")
	opts.SetUsername(os.Getenv("HIVE_USER"))
	opts.SetPassword(os.Getenv("HIVE_PSWD"))
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	name := "user" + strconv.Itoa(rand.IntN(100))
	password := "password" + strconv.Itoa(rand.IntN(100))
	age := rand.IntN(40)
	hours_spent := rand.IntN(100)
	
	datajson,_ :=  json.Marshal(map[string]interface{}{ "name": name, "password": password, "age": age, "hours_spent": hours_spent})
	token := client.Publish("test_topic/fazol", 1, false, datajson)
	token.Wait()
	
	time.Sleep(2 * time.Second)
	return Data{name, password, age, hours_spent}
}

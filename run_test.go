package main

import (
	"os"
	"testing"

	godotenv "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func TestPipeline(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Errorf("Error loading .env file: %s", err)
	}
	message := PublisherOne()
	
	mongo := Mongo(os.Getenv("mongopassword"))

	collection := GetCollection("test", "test", mongo)

	filter := bson.D{{Key: "name", Value: message.name}, {Key: "password", Value: message.password}, {Key: "age", Value: message.age}, {Key: "hours_spent", Value: message.hours_spent}}

	result := Select(collection, filter)

	if result.name != message.name {
		t.Errorf("Expected %s, got %s", message.name, result.name)
	}
	
	if result.password != message.password {
		t.Errorf("Expected %s, got %s", message.password, result.password)
	}

	if result.age != message.age {
		t.Errorf("Expected %d, got %d", message.age, result.age)
	}

	if result.hours_spent != message.hours_spent {
		t.Errorf("Expected %d, got %d", message.hours_spent, result.hours_spent)
	}

	t.Log("Test passed")
}
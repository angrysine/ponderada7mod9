package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	godotenv "github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func main() {
	err := godotenv.Load(".env")
	throw(err)
	var password = os.Getenv("mongopassword")
	fmt.Printf("password: %v \n",password)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//mongodb+srv://albertomiranda:sA8R8Ht5RgfDv2ed@betinhodb.xeezpin.mongodb.net/
	defer cancel()
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://albertomiranda:" + password + "@betinhodb.xeezpin.mongodb.net/?retryWrites=true&w=majority&appName=betinhoDb"))

	throw(err)
	TestConection(client, ctx)
	collection := GetCollection("betinhoDb","users",client)
	go Publisher()
	Subscriber(collection)
	select {}
}

func TestConection(client *mongo.Client, ctx context.Context) bool {
	err := client.Ping(ctx, readpref.Primary()) 

	if (err == nil) {

		return true
	}

	fmt.Printf("%v",err)

	return false
}

func GetCollection(databaseName string, collectionName string, client *mongo.Client) *mongo.Collection{
	return client.Database(databaseName).Collection(collectionName)
}


func Select(collection *mongo.Collection, filter bson.D) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var results []Data
	if err = cursor.All(ctx,results); err != nil {
		panic(err)
	}
	for _,result := range results {
		log.Println("name: "+result.name, "password: "+result.password, "age: "+strconv.Itoa(result.age), "hours_spent: "+strconv.Itoa(result.hours_spent))
	}
	
	
}

func Insert(collection *mongo.Collection, data Data) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, bson.D{{Key: "name",Value: data.name},{Key:"password",Value:data.password},{Key:"age",Value:data.age},{Key:"hours_spent",Value:data.hours_spent}})
	fmt.Printf("insert item with id %v",res.InsertedID)
	if err != nil {
		log.Fatal(err)
	}
}

func throw(err error) {
	if (err != nil) {
		panic(err)
	}
}



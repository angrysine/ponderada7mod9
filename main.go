package main

import (
	"math/rand/v2"
	"os/exec"
	"strconv"

	godotenv "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	err := godotenv.Load(".env")
	throw(err)
	go Publisher()
	go RunMetabase()
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

func RunMetabase() {
	cmd := exec.Command("docker", "compose", "up")
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
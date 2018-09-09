package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printObject(obj interface{}) {
	mapB, _ := json.Marshal(obj)
	fmt.Println(string(mapB))
}

func main() {
	errLoadEnv := godotenv.Load()
	handleError(errLoadEnv)
	TOKEN := os.Getenv("TOKEN")
	api := slack.New(TOKEN)
	userMap := FetchUserWithIM(api)
	for k, v := range userMap {
		log.Printf("Key: %s, ID: %s, Name: %s, ImID: %s\n", k, v.UserID, v.UserName, v.ImID)
	}

	usersLogWorkInformation := ReadDataFromFile("data/all.txt")
	for _, v := range usersLogWorkInformation {
		// printObject(v)
		log.Println(v.messageNotice())
		log.Println()
		log.Println()
	}

	SendMessageIntoListUser(api, usersLogWorkInformation, userMap)
	log.Println()
	log.Println()
}

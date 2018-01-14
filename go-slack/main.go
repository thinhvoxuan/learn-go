package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

type userWithIM struct {
	UserID   string
	UserName string
	ImID     string
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchuserWithIM(api *slack.Client) (userMap map[string]userWithIM) {
	users, errUsers := api.GetUsers()
	handleError(errUsers)

	userMapIDToValue := map[string]userWithIM{}
	for _, u := range users {
		if !u.IsBot {
			log.Printf("ID: %s, Name: %s\n", u.ID, u.Name)
			uwi := userWithIM{u.ID, u.Name, ""}
			userMapIDToValue[u.ID] = uwi
		}
	}

	ims, errIms := api.GetIMChannels()
	handleError(errIms)

	for _, im := range ims {
		uwi, ok := userMapIDToValue[im.User]
		if ok {
			log.Printf("ID: %s, Name: %s\n", im.ID, im.User)
			uwi.ImID = im.ID
			userMapIDToValue[im.User] = uwi
		}
	}

	userMap = map[string]userWithIM{}
	for _, v := range userMapIDToValue {
		log.Printf("ID: %s, Name: %s, ImID: %s\n", v.UserID, v.UserName, v.ImID)
		userMap[v.UserName] = v
	}
	return
}

func sendMessageIntoListUser(api *slack.Client, message string, listUsers []string, userMap map[string]userWithIM) {
	for _, userName := range listUsers {
		if user, ok := userMap[userName]; ok {
			params := slack.PostMessageParameters{AsUser: true}
			channelID, timestamp, errorMessage := api.PostMessage(user.ImID, message, params)
			handleError(errorMessage)
			log.Printf("Message successfully sent to user %s with channel %s at %s", user.UserName, channelID, timestamp)
		}
	}
}

func file2lines(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return lines
}

func main() {
	errLoadEnv := godotenv.Load()
	handleError(errLoadEnv)
	TOKEN := os.Getenv("TOKEN")
	api := slack.New(TOKEN)
	userMap := fetchuserWithIM(api)
	for k, v := range userMap {
		log.Printf("Key: %s, ID: %s, Name: %s, ImID: %s\n", k, v.UserID, v.UserName, v.ImID)
	}
	message := "Hello from API"
	usersCompleted := file2lines("data/completed.txt")
	sendMessageIntoListUser(api, message, usersCompleted, userMap)
}

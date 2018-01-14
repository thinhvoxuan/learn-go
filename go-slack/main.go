package main

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
)

type UserWithIM struct {
	UserID   string
	UserName string
	ImID     string
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func fetchUserWithIm(api *slack.Client) (userMap map[string]UserWithIM) {
	users, errUsers := api.GetUsers()
	handleError(errUsers)

	userMapIDToValue := map[string]UserWithIM{}
	for _, u := range users {
		if !u.IsBot {
			log.Printf("ID: %s, Name: %s\n", u.ID, u.Name)
			uwi := UserWithIM{u.ID, u.Name, ""}
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

	userMap = map[string]UserWithIM{}
	for _, v := range userMapIDToValue {
		log.Printf("ID: %s, Name: %s, ImID: %s\n", v.UserID, v.UserName, v.ImID)
		userMap[v.UserName] = v
	}
	return
}

func sendMessageIntoListUser(api *slack.Client, message string, listUsers []string, userMap map[string]UserWithIM) {
	for _, userName := range listUsers {
		if user, ok := userMap[userName]; ok {
			params := slack.PostMessageParameters{AsUser: true}
			channelID, timestamp, errorMessage := api.PostMessage(user.ImID, message, params)
			handleError(errorMessage)
			log.Printf("Message successfully sent to user %s with channel %s at %s", user.UserName, channelID, timestamp)
		}
	}
}

func main() {
	errLoadEnv := godotenv.Load()
	handleError(errLoadEnv)
	TOKEN := os.Getenv("TOKEN")
	USERLIST := os.Getenv("USERS")
	api := slack.New(TOKEN)
	userMap := fetchUserWithIm(api)
	for k, v := range userMap {
		log.Printf("Key: %s, ID: %s, Name: %s, ImID: %s\n", k, v.UserID, v.UserName, v.ImID)
	}
	message := "Hello from API"
	users := strings.Split(USERLIST, ";")
	sendMessageIntoListUser(api, message, users, userMap)
}

package main

import (
	"log"

	"github.com/nlopes/slack"
)

// UserWithIM mapping slack alias with ImID
type UserWithIM struct {
	UserID   string
	UserName string
	ImID     string
}

// FetchUserWithIM Fetch all user & IMs then return mapping
func FetchUserWithIM(api *slack.Client) (userMap map[string]UserWithIM) {
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

// SendMessageIntoListUser Send message Into List user
func SendMessageIntoListUser(api *slack.Client, listUsers []UserLogWork, userMap map[string]UserWithIM) {
	for _, user := range listUsers {
		if userInSlack, ok := userMap[user.UserAlias]; ok {
			params := slack.PostMessageParameters{AsUser: true, LinkNames: 1}
			message := user.messageNotice()
			log.Println(message)
			channelID, timestamp, errorMessage := api.PostMessage(userInSlack.ImID, message, params)
			handleError(errorMessage)
			log.Printf("Message successfully sent to user %s with channel %s at %s", userInSlack.UserName, channelID, timestamp)
		}
	}
}

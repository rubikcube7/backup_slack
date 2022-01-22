package main

import (
	"backup_slack/config"
	"fmt"
	"github.com/slack-go/slack"
)

type GetConversationHistoryParameters struct {
	ChannelID string
	Cursor    string
	Inclusive bool
	Latest    string
	Limit     int
	Oldest    string
}

func main() {

	api := slack.New(config.Config.ApiToken)
	user, err := api.GetUserInfo(config.Config.ApiUserID)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)

	params := GetConversationHistoryParameters{ChannelID: config.Config.ApiChannelID, Limit: 1000}
	res, err := api.GetConversationHistory((*slack.GetConversationHistoryParameters)(&params))

	for _, m := range res.Messages {
		//for _, att := range m.Attachments {
		fmt.Printf("Message: %s,\n", m.Text)
		//}
	}
}

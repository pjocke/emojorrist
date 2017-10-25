package main

import (
	"os"
	"log"
	"github.com/nlopes/slack"
)

func main() {
	key, ok := os.LookupEnv("SLACK_API_TOKEN")
	if !ok {
		log.Fatal("Please set SLACK_API_TOKEN environment variable.")
	}

	api := slack.New(key)
	rtm := api.NewRTM()

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			emoji := GenerateRandomEmoji()

			ref := slack.NewRefToMessage(ev.Channel, ev.Timestamp)
			err := api.AddReaction(emoji, ref)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

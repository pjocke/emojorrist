package main

import (
	"os"
	"log"
	"github.com/nlopes/slack"
	"fmt"
	"strings"
)

func main() {
	matches := map[string]string{
		"sko":		"mans_shoe",
		"skor":		"mans_shoe",
		"skjorta":	"shirt",
		"skjortor":	"shirt",
		"kostym":	"man_in_business_suit_levitating",
		"kostymer":	"man_in_business_suit_levitating",
		"bajs":		"poop",
		"skit":		"poop",
		"bög":		"two_men_holding_hands",
		"bögar":	"two_men_holding_hands",
		"indier":	"man_with_turban",
		"idiot":	"man_with_turban",
		"polis":	"cop",
		"polisen":	"cop",
		"poliser":	"cop",
		"svamp":	"mushroom",
		"spöke":	"ghost",
		"jul":		"christmas_tree",
		"tomte":	"santa",
		"slips":	"necktie",
		"slipsar":	"necktie",
		"öl":		"beers",
		"burgare":	"hamburger",
		"hamburgare":	"hamburger",
		"kuk":		"eggplant",
		"kuken":	"eggplant",
		"kukar":	"eggplant",
		"vin":		"whine_glass",
		"kaffe":	"coffee",
		"cocktail":	"cocktail",
		"cocktails":	"cocktail",
		"pommes":	"fries",
		"pizza":	"pizza",
		"skjut":	"gun",
		"knark":	"pill",
		"knarka":	"pill",
		"pengar":	"moneybag",
		"dusch":	"shower",
		"duscha":	"shower",
		"flyg":		"airplane",
		"flyga":	"airplane",
		"flyget":	"airplane",
		"hallonsorbet":	"shaved_ice",
		"röv":		"peach",
		"häst":		"horse",
		"hästen":	"horse",
		"hästar":	"horse",
		"kaka":		"cookie",
		"tårta":	"birthday",
		"jennie":	"weary",
	}

	key, ok := os.LookupEnv("SLACK_API_TOKEN")
	if !ok {
		log.Fatal("Please set SLACK_API_TOKEN environment variable.")
	}

	api := slack.New(key)
	rtm := api.NewRTM()

	logger := log.New(os.Stdout, "emoji: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			fmt.Printf("%s\n", ev.Text)
			ref := slack.NewRefToMessage(ev.Channel, ev.Timestamp)

			words := strings.Fields(ev.Text)
			for _, word := range words {
				word = strings.Trim(strings.ToLower(word), ".,:;!?")
				if emoji, ok := matches[word]; ok {
					fmt.Printf("match: %s = %s\n", word, emoji)
					err := api.AddReaction(emoji, ref)
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	}

}

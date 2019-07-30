package main

import (
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"github.com/nlopes/slack"
	"time"

	"./util"
)

var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))

func main()  {
	log.Println("Starting up ChenBot...")
	botConfiguration := util.ReadConfigFile("./config")

	// Setup slack client and
	slackClient := slack.New(botConfiguration.OAuthAccessToken)
	slackRTM := slackClient.NewRTM()
	// Handle all communications in another thread
	go slackRTM.ManageConnection()

	messageHandler:
		for {
			select {
			case msg := <- slackRTM.IncomingEvents:
				fmt.Println("Event received: " + msg.Type)
				switch event := msg.Data.(type) {
				case *slack.MessageEvent:
					info := slackRTM.GetInfo()

					// check, is the message sender the bot itself? Prevents self calling
					if event.User == info.User.ID {
						continue messageHandler // skip over, do nothing
					}

					if event.BotID != "" {
						fmt.Println("Bot ID: ", event.BotID)
					}

					text := event.Text
					text = strings.TrimSpace(text)
					text = strings.ToLower(text)

					// pass of to handler; uses regex to determine appropriate
					// TODO: Implement this
					go handleMessage(text, event, slackRTM)

				case *slack.RTMError:
					log.Printf("Error: %s\n", event.Error())

				case *slack.InvalidAuthEvent:
					log.Fatalf("Error: Invalid credentials")

				default:
					// do nothing

					log.Println(event)
				}
			}
		}
}

func handleMessage(message string, event *slack.MessageEvent, slackRTM *slack.RTM) {

	easyButtonRegexp, err := regexp.Compile("(.*(e+a+[sz]+y+|e+z+|e+ +z+).*)")
	if err != nil {
		log.Fatalf("Error generating easyButtonRegexp: %s", err)
	}

	if easyButtonRegexp.MatchString(message) {
		if event.User == "UL1MWS8D6" && randGenerator.Int() % 4 == 0 {
			slackRTM.SendMessage(slackRTM.NewOutgoingMessage("No Chen, YOU'RE EZ :easy:", event.Channel))
			return
		}

		slackRTM.SendMessage(slackRTM.NewOutgoingMessage(":easy:", event.Channel))
	}

	jamRegexp, err := regexp.Compile(".*@jam.*")
	if err != nil {
		log.Fatalf("Error generating jamRegexp: %s", err)
	}

	if jamRegexp.MatchString(message) {
		slackRTM.SendMessage(slackRTM.NewOutgoingMessage("<@UKQA8VBHR>, <@UL5194VE2>", event.Channel))
	}

	chamRegexp, err := regexp.Compile(".*@chyam.*")
	if err != nil {
		log.Fatalf("Error generating chamRegexp: %s", err)
	}

	if chamRegexp.MatchString(message) {
		slackRTM.SendMessage(slackRTM.NewOutgoingMessage("<@UL1MWS8D6>, <@UL3HFQPHD>", event.Channel))
	}

	deringRegexp, err := regexp.Compile(".*@dering.*")
	if err != nil {
		log.Fatalf("Error generating deringRegexp: %s", err)
	}

	if deringRegexp.MatchString(message) {
		slackRTM.SendMessage(slackRTM.NewOutgoingMessage("<@UKQBVGZGT>, <@UKQ9P39B4>", event.Channel))
	}

	shrevorRegexp, err := regexp.Compile(".*@shrevor.*")
	if err != nil {
		log.Fatalf("Error generating shrevorRegexp: %s", err)
	}

	if shrevorRegexp.MatchString(message) {
		slackRTM.SendMessage(slackRTM.NewOutgoingMessage("<@ULJTDFB4H>, <@UL3HFQPHD>", event.Channel))
	}


}



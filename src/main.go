package main

import (
	"./util"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type GroupAlias struct {
	RegexString string
	MentionString string
}

var randGenerator = rand.New(rand.NewSource(time.Now().UnixNano()))
var personMap = util.IntializeMapping()
var groupAliases []GroupAlias = []GroupAlias{
	{RegexString: ".*@jam.*", MentionString: fmt.Sprintf("<@%s>, <@%s>", personMap["Jenna"], personMap["Sam"])},
	{RegexString: ".*@chyam.*", MentionString: fmt.Sprintf("<@%s>, <@%s>", personMap["Chen"], personMap["Shyam"])},
	{RegexString: ".*@dering.*", MentionString: fmt.Sprintf("<@%s>, <@%s>", personMap["Derek"], personMap["Wenting"])},
	{RegexString: ".*@shrevor.*", MentionString: fmt.Sprintf("<@%s>, <@%s>", personMap["Shyam"], personMap["Trevor"])},
}

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

	for _, groupAlias := range groupAliases {
		groupAlias.handleGroupAlias(message, event, slackRTM)
	}
}

func (ga GroupAlias) handleGroupAlias(message string, event *slack.MessageEvent, slackRTM *slack.RTM) {
	groupAliasRegex, err := regexp.Compile(ga.RegexString)
	if err != nil {
		log.Fatalf("Error generating expression %s with: %s", ga.RegexString, err)
	}

	if groupAliasRegex.MatchString(message) {
		slackRTM.SendMessage(slackRTM.NewOutgoingMessage(ga.MentionString, event.Channel))
	}
}

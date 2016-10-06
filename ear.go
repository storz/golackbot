package golackbot

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/nlopes/slack"
)

func (bot *golackbot) Listen() {
	enc := json.NewEncoder(os.Stdout)
	rtm := bot.api.NewRTM()
	go rtm.ManageConnection()

	color.Cyan("Ready")

Loop:
	for {
		select {
		case ie := <-rtm.IncomingEvents:
			switch ev := ie.Data.(type) {
			case *slack.MessageEvent:
				fmt.Print("[Message] ")
				msg := ev.Msg
				enc.Encode(msg)
				bot.process(msg)

			case *slack.RTMError:
				color.Red("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				color.Red("Invalid credentials")
				break Loop

			default:

			}
		}
	}
}

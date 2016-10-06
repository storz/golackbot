package golackbot

import (
	"github.com/fatih/color"
	"github.com/nlopes/slack"
)

func (bot *golackbot) SendTo(msg slack.Msg, text string, channel string) {
	params := slack.PostMessageParameters{}
	params.AsUser = true
	params.LinkNames = 1
	params.UnfurlMedia = true

	channelID, timestamp, err := bot.api.PostMessage(channel, text, params)
	if err != nil {
		color.Red("%s\n", err)
		return
	}
	color.Green(bot.SayMyName()+" says: '%s' to channel %s at %s\n", text, channelID, timestamp)
}

func (bot *golackbot) Say(msg slack.Msg, text string) {
	bot.SendTo(msg, text, msg.Channel)
}

func (bot *golackbot) Reply(msg slack.Msg, text string) {
	text = "<@" + msg.User + ">: " + text
	bot.Say(msg, text)
}

package golackbot

import "github.com/nlopes/slack"

// GolackBot has commands like a Hubot
type GolackBot interface {
	Listen()
	Respond(reg string, f func(msg slack.Msg))
	Hear(reg string, f func(msg slack.Msg))
	SendTo(msg slack.Msg, text string, channel string)
	Say(msg slack.Msg, text string)
	Reply(msg slack.Msg, text string)
	LoadConfig(path string)
	SayMyName() string
	GenerateMsgLink(msg slack.Msg) string
}

type golackbot struct {
	token            string
	id               string
	name             string
	aka              []string
	api              *slack.Client
	respondReactions []reaction
	hearReactions    []reaction
}

type reaction struct {
	reg  string
	exec func(slack.Msg)
}

// NewBot is initializer. But instance has no identity
func NewBot() GolackBot {
	return &golackbot{}
}

func (bot *golackbot) SayMyName() string {
	return bot.name
}

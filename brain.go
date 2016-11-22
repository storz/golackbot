package golackbot

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/nlopes/slack"
)

func match(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func extract(reg, str string) string {
	return regexp.MustCompile(reg).ReplaceAllString(str, "")
}

func (bot *golackbot) getMyReg(regstr string) string {
	return "(?i)^@?\u003c?(" + strings.Join(append(bot.aka, bot.name, bot.id)[:], "|") + ")\u003e?:?(\\s+)+" + extract(`^\^`, regstr)
}

func (bot *golackbot) process(msg slack.Msg) {
	if msg.SubType == "bot_message" && msg.Username == bot.name {
		return
	}
	text := msg.Text

	for _, r := range bot.respondReactions {
		if match(bot.getMyReg(r.reg), text) || r.reg == "" {
			r.exec(msg)
			if r.reg != "" {
				break
			}
		}
	}
	for _, r := range bot.hearReactions {
		if match(r.reg, text) || r.reg == "" {
			r.exec(msg)
			if r.reg != "" {
				break
			}
		}
	}
}

func (bot *golackbot) Respond(reg string, f func(msg slack.Msg)) {
	bot.respondReactions = append(bot.respondReactions, reaction{reg: reg, exec: f})
	fmt.Printf("Responding item is added: %s\n", reg)
}

func (bot *golackbot) Hear(reg string, f func(msg slack.Msg)) {
	bot.hearReactions = append(bot.hearReactions, reaction{reg: reg, exec: f})
	fmt.Printf("Hearing item is added: %s\n", reg)
}

func (bot *golackbot) GenerateMsgLink(msg slack.Msg) string {
	teamInfo, err := bot.api.GetTeamInfo()
	if err != nil {
		panic(err)
	}
	teamDomain := teamInfo.Domain

	var chOrUser string
	channelInfo, err := bot.api.GetChannelInfo(msg.Channel)
	if err != nil {
		if err.Error() == "channel_not_found" {
			chOrUser = msg.Channel
		} else {
			panic(err)
		}
	} else {
		chOrUser = channelInfo.Name
	}

	ts := extract(`\D`, msg.Timestamp)
	return fmt.Sprintf("https://%s.slack.com/archives/%s/p%s", teamDomain, chOrUser, ts)
}

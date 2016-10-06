package golackbot

import (
	"github.com/BurntSushi/toml"
	"github.com/nlopes/slack"
)

type config struct {
	Bot botConfig `toml:"bot"`
}

type botConfig struct {
	Token string   `toml:"token"`
	ID    string   `toml:"id"`
	Name  string   `toml:"name"`
	Aka   []string `toml:"aka"`
}

func (bot *golackbot) LoadConfig(path string) {
	var cfg config
	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		panic(err)
	}
	bot.id = cfg.Bot.ID
	bot.name = cfg.Bot.Name
	bot.aka = cfg.Bot.Aka
	bot.api = slack.New(cfg.Bot.Token)
}

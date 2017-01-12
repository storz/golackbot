# GolackBot

GolackBot is simple slack bot like a Hubot.

## Usage

### Put config file

```Toml
# config.tml
[bot]
token    = "xoxb-XXXXXXXXXXXXXXXXX"
id       = "@XXXXXXX"
name     = "Kan"
aka      = [ "GAMI" ]
```

### Create main.go

```Go
// main.go
package main

import (
	"path/filepath"

	"github.com/storz/golackbot"
)

var (
	bot = golackbot.NewBot()
)

func main() {
	cfgPath, err := filepath.Abs("config.tml")
	if err != nil {
		panic(err)
	}
	bot.LoadConfig(cfgPath)
	bot.Listen()
}
```

### Scripting

```Go
// goodcommunication.go
package main

import "github.com/nlopes/slack"

func init() {
    bot.Respond("ping", func(msg slack.Msg) { // React to mention
        // Some processes...
        // bot.Reply(msg, "pong") // Kan says "@your_name: pong"
    })

    // You can use regexp on 1st arg
    bot.Hear("^Yo$", func(msg, slack.Msg) { // React to any messages
        // Some processes...
        // bot.Say(msg, "pong")  // Kan says "Yo"
    }
}
```

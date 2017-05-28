package main

import (
	"fmt"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type yesBot struct {
	bot reddit.Bot
}

func (r *yesBot) Comment(c *reddit.Comment) error {
	fmt.Println(c.Body)
	return nil
}

func main() {

	if bot, err := reddit.NewBotFromAgentFile(".agent", 0); err != nil {
		fmt.Println("Failed to create bot handle: ", err)
	} else {
		cfg := graw.Config{SubredditComments: []string{"i_just_say_yes"}}
		handler := &yesBot{bot: bot}
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
	}
}

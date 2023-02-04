package main

import (
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {

}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xapp-1-A04LU9DR08N-4717389981172-489562202d3307af59fdacac03ede4867726cac7a327e770fd70f30c2f41f34e")
	os.Setenv("SLACK_APP_TOKEN", "xoxb-4679221292385-4738625359232-vxpmPVpx0PuzGnobymXC4hqQ")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
}

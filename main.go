package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jcksnvllxr80/slack-age-bot/secrets"
	"github.com/shomali11/slacker"
)

func printCommandEvents(testChannel <-chan *slacker.CommandEvent) {
	for event := range testChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	secrets.SetEnvVars()
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("my dob is <date>", &slacker.CommandDefinition{
		Description: "dob calculator",
		Example:     "my dob is 2022",
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			date := request.Param("date")
			layout := "2006-01-02"
			t, err := time.Parse(layout, date)
			if err != nil {
				fmt.Println("Error while parsing date :", err)
			}
			birthYear := t.Year()
			birthYearday := t.YearDay()
			currentTime := time.Now()
			currentYear := currentTime.Year()
			currentYearday := currentTime.YearDay()
			age := currentYear - birthYear
			fmt.Println("current yearday: ", currentYearday)
			fmt.Println("birth yearday: ", birthYearday)
			if currentYearday < birthYearday {
				fmt.Println("no birthday yet this year. So, youre not ", age)
				age = age - 1
				fmt.Println("Rather, you are ", age)
			}
			r := fmt.Sprintf("age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

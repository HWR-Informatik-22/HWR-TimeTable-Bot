package main

import (
	"hwrbot/helper"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

const course = "kursa"

func main() {
	godotenv.Load(".env")

	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	discord.Open()

	s := gocron.NewScheduler(time.UTC)
	s.Cron(os.Getenv("EXECUTE_INTERVAL")).Do(sendTimeTable, discord, "kursa", "kursb")
	s.StartBlocking()
}

func sendTimeTable(discord *discordgo.Session, courses ...string) {
	for _, course := range courses {
		content, err := helper.GetIcsContent(course)
		if err != nil {
			panic(err)
		}

		lessons, err := helper.ParseIcsContent(content)
		if err != nil {
			panic(err)
		}

		embed := helper.CreateEmbed(lessons, course)
		_, err = discord.ChannelMessageSendEmbed(os.Getenv("DISCORD_CHANNEL_ID"), embed)
		if err != nil {
			panic(err)
		}
	}
}

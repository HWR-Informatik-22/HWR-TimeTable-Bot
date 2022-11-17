package main

import (
	"fmt"
	"hwrbot/helper"
	"os"

	"github.com/bwmarrin/discordgo"
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

	content, err := helper.GetIcsContent(course)
	if err != nil {
		panic(err)
	}

	lessons, err := helper.ParseIcsContent(content)
	if err != nil {
		os.Exit(2)
		return
	}

	embed := helper.CreateEmbed(lessons, course)
	_, err = discord.ChannelMessageSendEmbed(os.Getenv("DISCORD_CHANNEL_ID"), embed)
	if err != nil {
		panic(err)
	}

	fmt.Scanln()
}

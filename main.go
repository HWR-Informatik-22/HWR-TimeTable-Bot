package main

import (
	"hwrbot/helper"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const course = "kursa"

func main() {
	godotenv.Load(".env")
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	//discord.Channel(os.Getenv("DISCORD_CHANNEL_ID"))

	if err != nil {
		panic(err)
	}

	discord.Open()

	godotenv.Load(".env")
	file, err := helper.GetIcsFile(course)

	if err != nil {
		panic(err)
	}

	lessons := helper.ParseIcsFile(file)
	embed := helper.CreateEmbed(lessons, course)

	_ = embed

	//_, err = discord.ChannelMessageSendEmbed(os.Getenv("DISCORD_CHANNEL_ID"), embed)
	if err != nil {
		panic(err)
	}

	//fmt.Scanln()
}

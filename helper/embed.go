package helper

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func CreateEmbed(lessons *[]Lesson, course string) *discordgo.MessageEmbed {
	return &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s - %s", time.Now().Format("02.01.2006"), course),
		Color:  0xde1d10,
		Fields: getEmbedFields(lessons, course),
	}
}

func getEmbedFields(lessons *[]Lesson, course string) []*discordgo.MessageEmbedField {
	fields := make([]*discordgo.MessageEmbedField, len(*lessons))
	for i, e := range *lessons {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s (%s)", e.name, e.teacher),
			Value:  fmt.Sprintf("Von: %s\nBis: %s\n%s", e.start, e.end, e.room),
			Inline: false,
		}
	}
	return fields
}

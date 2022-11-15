package helper

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func createEmbed(lessons []Lesson, course string) {
	x := discordgo.MessageEmbed{
		Title:  "",
		Color:  0xde1d10,
		Fields: getEmbedFields(lessons, course),
	}

	_ = x
}

func getEmbedFields(lessons []Lesson, course string) []*discordgo.MessageEmbedField {
	fields := make([]*discordgo.MessageEmbedField, len(lessons))
	for i, e := range lessons {
		fields[i] = &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s (%s)", e.name, e.teacher),
			Value:  fmt.Sprintf("%s\n%s\n%s", e.start, e.end, e.room),
			Inline: false,
		}
	}
	return fields
}

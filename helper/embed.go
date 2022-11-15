package helper

import "github.com/bwmarrin/discordgo"

func createEmbed(lesson Lesson) {
	x := discordgo.MessageEmbed {
		Title: lesson.name,
		Color: 0xde1d10,
		Fields: ,
	}

	_ = x;
}

func getEmbedFields(lessons []Lesson, course string)[]*discordgo.MessageEmbedField{
	fields := make([]*discordgo.MessageEmbedField, len(lessons))
	return 
}
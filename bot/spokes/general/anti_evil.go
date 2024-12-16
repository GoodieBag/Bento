package general

import (
	"main/bot"
	"main/bot/spokes/evil"

	"github.com/bwmarrin/discordgo"
)

type AntiEvil struct{}

func GetAntiEvil() *AntiEvil {
	return &AntiEvil{}
}

func (p *AntiEvil) Commands() bot.BotCommandMap {
	cmdMap := make(bot.BotCommandMap)

	cmdMap["status"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "")
	}
	return cmdMap
}

func (p *AntiEvil) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Author.ID == "1276413143299522685" && m.Content == evil.FightPhrase {
		s.ChannelMessageSend(m.ChannelID, "<@1276413143299522685> standdown")
		s.ChannelMessageSend(m.ChannelID, "The best fights are the ones you stay out of.")
	}
}

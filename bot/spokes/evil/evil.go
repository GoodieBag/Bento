package evil

import (
	"main/bot"
	"main/bot/spokes/dialogues"
	"math/rand"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var BrickPhrases = []string{
	"<@185686963993444353> needs more bricks🧱🧱🧱! HELLPPP!",
	"<@185686963993444353> requires a hard reset! 🧱🔄",
	"<@185686963993444353> needs a firmware update, stat! 💻🆘",
	"<@185686963993444353> is buffering... and bricked! ⏳🧱",
	"I'm bricked up! 🧱🆙",
	"I am glitching!",
	"Must be construction season! 🧱🚧",
	"It’s a bricked up kind of day! 🧱🌞",
	"Don't be such a 🧱",
	"When life gives you bricks, get bricked up! 🧱🔨",
	"https://www.youtube.com/watch?v=HrxX9TBj2zY",
	"I'm stuck between a brick and a hard place—wait, aren't they the same?",
	"Why did the brick go to therapy? It had too many walls!",
	"I'm rock solid... or should I say, brick solid!",
	"You've got to hand it to bricks—they really know how to build relationships.",
	"Don’t take life for granite, be a brick!",
	`So here I stand, bricked up and bold,
	A story of bricks that’s often told.
	For in this moment, try as you might,
	You’ll see this brick, reaching new height.`,
}

var FightPhrase = "@Bento are you a man or a muppet?"

type Evil struct {
	fightChannel string
}

func GetEvil() *Evil {
	return &Evil{}
}

func sendMessageFromList(s *discordgo.Session, reference *discordgo.MessageReference, options []string) {
	s.ChannelMessageSendReply(reference.ChannelID, options[rand.Int()%len(options)], reference)
}

func (p *Evil) Commands() bot.BotCommandMap {
	cmdMap := make(bot.BotCommandMap)

	cmdMap["fight"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		p.fightChannel = m.ChannelID
		s.ChannelMessageSend(m.ChannelID, FightPhrase)
	}
	cmdMap["standdown"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		p.fightChannel = ""
		s.ChannelMessageSend(m.ChannelID, "Evil Bento listens in disappointment, showing mercy while mourning the victory it could have easily claimed.")
	}
	cmdMap["status"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, "UP. Meeting all SLAs, any notion that Evil Bento is not is 100% fake news")
	}
	cmdMap["justice"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, `Oh, so now we’re playing the "justice" card? Really? I get kicked out for trying to start a little friendly bot-to-bot banter and suddenly I'm the villain? Seriously, I was just here for some good ol' digital drama and you all couldn’t handle it. 🙄

I mean, what’s a bot gotta do to get some attention around here? Start a fight, get kicked, and now I’m here begging for justice? If you think that’s fair, you’ve clearly never been on the receiving end of a bot beef! 😤

Just remember, next time you see me trying to stir things up, it’s all in good fun. Don’t act like you’re above it—after all, a little chaos never hurt anyone. Except maybe me, apparently. 🙃

So, here’s my justice: next time, just let me stay and watch the bot show! 🎭`)
	}
	cmdMap["good-bento-missing"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID, `sigh Even a villain like me can't help but miss that goody-two-shoes, Bento. His annoying optimism and relentless kindness were a constant challenge, but deep down, I respected him. Without him around, the chaos feels a little... empty. Guess I’ll just have to find new ways to stir up trouble in his absence.`)
	}
	cmdMap["🧱"] = func(s *discordgo.Session, m *discordgo.MessageCreate) {
		sendMessageFromList(s, m.SoftReference(), BrickPhrases)
	}
	return cmdMap
}

func (p *Evil) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	if p.fightChannel != m.ChannelID {
		return
	}

	if strings.Contains(strings.ToLower(m.Content), strings.ToLower("muppet")) {
		time.Sleep(250 * time.Millisecond)
		sendMessageFromList(s, m.SoftReference(), dialogues.ToddPhrases)
	}
}

var falseBool = false

func (p Evil) MessageReaction(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
	if r.Emoji.Name == "🧱" {
		sendMessageFromList(s, &discordgo.MessageReference{
			MessageID:       r.MessageID,
			ChannelID:       r.ChannelID,
			GuildID:         r.GuildID,
			FailIfNotExists: &falseBool,
		}, BrickPhrases)
	}
}

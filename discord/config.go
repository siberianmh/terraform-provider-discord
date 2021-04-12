package discord

import "github.com/bwmarrin/discordgo"

// Config exports the the needed to use properties
type Config struct {
	APIToken string
	GuildID  string
	Session  *discordgo.Session
}

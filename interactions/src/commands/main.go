package commands

import "github.com/bwmarrin/discordgo"

type InteractionHandler interface {
	Handle(*discordgo.Interaction) *discordgo.InteractionResponse
}

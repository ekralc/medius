package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var s *discordgo.Session

var AppID = "1079541448820662412"
var GuildID = "1079542109482258483"

func BoolAddr(value bool) *bool {
	val := value
	return &val
}

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "To test the bot's responsiveness.",
	},
	{
		Name: "Caption",
		Type: discordgo.MessageApplicationCommand,
	},
	{
		Name:         "Convert WebM to MP4",
		Type:         discordgo.MessageApplicationCommand,
		DMPermission: BoolAddr(true),
	},
}

func main() {
	s, _ = discordgo.New("Bot " + BotToken)

	fmt.Println("Overwriting local commands.")
	createdCommands, err := s.ApplicationCommandBulkOverwrite(AppID, "", commands)

	fmt.Println("\nCreated commands:")
	for _, cmd := range createdCommands {
		fmt.Println(cmd.Name)
	}

	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully created %v commands\n", len(createdCommands))
}

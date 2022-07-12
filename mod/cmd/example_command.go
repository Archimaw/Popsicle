package cmd

import (
	"bot/popsicle/v1/mod/uti"

	"github.com/bwmarrin/discordgo"
)

func init() {
	command := AddCommand("example", "An example command")
	command.SetHandler(
		func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			uti.NewEmbed().
				SetTitle("Example").
				SetDescription("An example command").
				SendInteractionMessage(s, i)
		},
	)
}

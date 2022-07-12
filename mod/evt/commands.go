package evt

import (
	"bot/popsicle/v1/mod/cmd"
	"log"

	"github.com/bwmarrin/discordgo"
)

func init() {
	AddEventHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if command, ok := cmd.CommandMap[i.ApplicationCommandData().Name]; ok {
			command.Handle(s, i)
			defer log.Printf("Command %s handled", i.ApplicationCommandData().Name)
		}
	})

	AddEventHandler(func(s *discordgo.Session, i *discordgo.Ready) {
		cmd.CreateAndUpdateCommands(s)
	})
}

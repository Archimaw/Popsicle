package cmd

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Command struct {
	Command *discordgo.ApplicationCommand
	Handle  func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

var CommandMap = make(map[string]*Command)

func AddCommand(name, desc string) *Command {
	cmd := &Command{
		Command: &discordgo.ApplicationCommand{
			Name:        name,
			Description: desc,
		},
	}
	CommandMap[name] = cmd
	return cmd
}

type CommandOption struct {
	Name, Description string
}

type CommandChoiceOption struct {
	*CommandOption
	Value interface{}
}

func (c *Command) AddOption(name, description string, opttype discordgo.ApplicationCommandOptionType) {
	c.Command.Options = append(c.Command.Options, &discordgo.ApplicationCommandOption{
		Name:        name,
		Description: description,
		Type:        opttype,
		Required:    true,
	})
}

func (c *Command) AddOptionChoices(name, desc string, optionType discordgo.ApplicationCommandOptionType, choices []CommandChoiceOption) {
	commandChoices := make([]*discordgo.ApplicationCommandOptionChoice, len(choices))
	for idx, choice := range choices {
		commandChoices[idx] = &discordgo.ApplicationCommandOptionChoice{
			Name:  choice.Name,
			Value: choice.Value,
		}
	}
	c.Command.Options = append(c.Command.Options, &discordgo.ApplicationCommandOption{
		Name:        name,
		Description: desc,
		Type:        optionType,
		Required:    false,
		Choices:     commandChoices,
	})
}

func (command *Command) SetHandler(handler func(*discordgo.Session, *discordgo.InteractionCreate)) {
	command.Handle = handler
}

func CreateAndUpdateCommands(s *discordgo.Session) {
	var commands []*discordgo.ApplicationCommand
	for _, command := range CommandMap {
		commands = append(commands, command.Command)
	}
	for _, guild := range s.State.Guilds {
		s.ApplicationCommandBulkOverwrite(s.State.User.ID, guild.ID, commands)
	}
	log.Printf("Updated commands for %d guilds", len(s.State.Guilds))
}

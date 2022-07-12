package evt

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type EventHandler struct {
	Handler interface{}
}

var EventHandlers []*EventHandler

func AddEventHandler(handler interface{}) {
	EventHandlers = append(EventHandlers, &EventHandler{
		Handler: handler,
	})
}

func HandleAllEvents(s *discordgo.Session) {
	for _, handler := range EventHandlers {
		s.AddHandler(handler.Handler)
	}
	log.Printf("Registered %d event handlers", len(EventHandlers))
}

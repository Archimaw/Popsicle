package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bot/popsicle/v1/mod/evt"

	"github.com/bwmarrin/discordgo"
)

var token string

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
	if token == "" {
		log.Fatal("Token not set", " -t <token>")
	}
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}
	evt.HandleAllEvents(dg)
	if err := dg.Open(); err != nil {
		log.Fatalf("error opening connection: %v", err)
		return
	}
	sc := make(chan os.Signal, 1)
	log.Println("Bot is now running. Press CTRL-C to exit.")
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
	dg.Close()
	log.Println("Bot is now closing.")
}

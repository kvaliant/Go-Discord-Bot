package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/go-discord-bot/config"
	"github.com/kvaliant/go-discord-bot/handler"
	"github.com/kvaliant/go-discord-bot/service"
)

type Config struct {
	Token string
	Application string
}

func init(){
	config.LoadEnvVariables()
}

func initBot(cfg *Config, callback func(ds *discordgo.Session)) {
	dc, err := discordgo.New("Bot "+ cfg.Token)
	if err != nil {
		log.Fatalf("Error while initiating bot, %v", err)
	}

	dc.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentGuildMessages | discordgo.IntentGuildMessageReactions | discordgo.IntentGuildMembers

	callback(dc)

	err = dc.Open()
	if err != nil { 
		log.Fatalf("Error while connecting websocket to discord server, %v", err)
	}

	defer dc.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Printf("Discord server websocket closed")
}

func main(){
	cfg := Config{}
	cfg.Token = os.Getenv("TOKEN")
	cfg.Application = os.Getenv("APPLICATION")

	fmt.Printf("Dc Token = "+cfg.Token)
	initBot(&cfg, func(ds *discordgo.Session){
		log.Printf("Discord server connection established")

		service := service.NewBotService()
		hdl := handler.NewBotHandler(service)
		ds.AddHandler(hdl.OnReady)
		ds.AddHandler(hdl.OnReactionAdd)
		ds.AddHandler(hdl.OnMessageCreate)
		ds.AddHandler(hdl.OnInteraction)
	})
}
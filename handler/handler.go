package handler

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kvaliant/go-discord-bot/service"
)

type handler struct {
	service service.Service
}

func NewBotHandler(svc service.Service) handler {
	return handler{service: svc}
}

func (h *handler) OnReady(session *discordgo.Session, message *discordgo.Ready){
	log.Println("Discord bot active, Listening for events")
	h.service.SendMessageToAllServers(session, "Golang Discord bot active")
}

func (h *handler) OnReactionAdd(session *discordgo.Session, e *discordgo.MessageReactionAdd){
	h.service.LogReactionAdd(session, e)
}

func (h *handler) OnMessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate){
	h.service.LogMessageCreate(session, msg)
}

func (h *handler) OnInteraction(session *discordgo.Session, in *discordgo.InteractionCreate){
	if in.Type == discordgo.InteractionApplicationCommand {
		command := in.ApplicationCommandData()
		switch commandName := command.Name; commandName {
		case "send_button":
			h.service.SendButton(session, in)
		case "send_text_input":
			h.service.SendTextInput(session, in)
		default:
			h.service.LogApplicationCommand(session, in)
		}
	}
	if in.Type == discordgo.InteractionModalSubmit {
		h.service.LogModalSubmit(session, in)
	}
}
package service

import (
	"log"

	"github.com/bwmarrin/discordgo"
)
func (s *Service) SendButton(session *discordgo.Session, in *discordgo.InteractionCreate){
	command := in.ApplicationCommandData()
	buttonType := ""
	if len(command.Options) != 0 {
		buttonType = command.Options[0].Value.(string)
	}

	button := &discordgo.Button{}
	switch buttonType {
	case "url":
		button.Label = "Go to Google"
		button.Style = discordgo.LinkButton
		button.URL = "http://google.com"
	case "success":
		button.Label = "A success button"
		button.Style = discordgo.SuccessButton
		button.CustomID = "-"
	default:
		button.Label = "A primary button"
		button.Style = discordgo.PrimaryButton
		button.CustomID = "-"
	}

	data := &discordgo.InteractionResponseData{}
	data.Content = "Here is a button"
	data.Components = []discordgo.MessageComponent{
		&discordgo.ActionsRow{Components: []discordgo.MessageComponent{
			*button,
		}},
	}

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: data,
	}
	
	err := session.InteractionRespond(in.Interaction, &response)
	if err != nil {
		log.Panicf("Err while interaction respond, %v", err)
	}
}
package service

import (
	"log"

	"github.com/bwmarrin/discordgo"
)
func (s *Service) SendTextInput(session *discordgo.Session, in *discordgo.InteractionCreate){
	textInput := &discordgo.TextInput{}
	textInput.CustomID = "This is a text input"
	textInput.Style = discordgo.TextInputStyle(2)
	textInput.Label = "This is a text input"

	textInput.Placeholder = "Text input placeholder"
	
	data := &discordgo.InteractionResponseData{}
	data.Content = "Here is a tect input inside a modal"
	data.CustomID = "This is a modal"
	data.Title = "Modal sample"
	data.Components = []discordgo.MessageComponent{
		&discordgo.ActionsRow{Components: []discordgo.MessageComponent{
			*textInput,
		}},
	}

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: data,
	}
	
	err := session.InteractionRespond(in.Interaction, &response)
	if err != nil {
		log.Panicf("Err while interaction respond, %v", err)
	}
}
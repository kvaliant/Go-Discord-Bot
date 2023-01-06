package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) LogApplicationCommand(session *discordgo.Session, in *discordgo.InteractionCreate){
	command := in.ApplicationCommandData()
	
	content := fmt.Sprintf("Command detected, name : %v , params : `%v`", command.Name, command.Options)
	for _, option := range command.Options {
		content = fmt.Sprintf("%v \n `Name : %v , value : %v` ", content, option.Name, option.Value)
	}

	log.Printf(content)
	
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			Content: content,
		},
	}
	
	err := session.InteractionRespond(in.Interaction, &data)
	if err != nil {
		log.Panicf("Err while interaction respond, %v", err)
	}
}
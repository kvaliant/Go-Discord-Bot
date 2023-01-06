package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) LogModalSubmit(session *discordgo.Session, in *discordgo.InteractionCreate) {
	modalData := in.Interaction.ModalSubmitData().Components[0]

	byteJson, _ := modalData.MarshalJSON()
	ar := &discordgo.ActionsRow{}
	err := ar.UnmarshalJSON(byteJson)
	if err != nil {
		log.Panicf("Err while unmarshal JSON, %v", err)
	}

	textInput := ar.Components[0]
	byteJson, _ = textInput.MarshalJSON()
	ti := &discordgo.TextInput{}
	err = discordgo.Unmarshal(byteJson, &ti)
	if err != nil {
		log.Panicf("Err while unmarshal JSON, %v", err)
	}

	content := fmt.Sprintf("Value received: %v", ti.Value)

	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	}

	err = session.InteractionRespond(in.Interaction, &data)
	if err != nil {
		log.Panicf("Err while interaction respond, %v", err)
	}
}
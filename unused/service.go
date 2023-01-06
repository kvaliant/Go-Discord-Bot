package unused

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	ChannelID string
}

func NewBotService(ChannelID string) Service {
	return Service{ChannelID: ChannelID}
}

func (s *Service) SendMessageToChannel(session *discordgo.Session, content string){
	data := discordgo.MessageSend{
		Content: content,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Style: discordgo.PrimaryButton,
						Label: "Click",
						CustomID: "-",
					},
				},
			},
		},
	}

	_, err := session.ChannelMessageSendComplex(s.ChannelID, &data)
	// _, err = session.ChannelMessageSend(ch.ID, fmt.Sprintf("Message"))

	if err != nil {
		log.Panicf("Error while channel message send complex, %v", err)
	}
}

func (s *Service) RespondInteraction(session *discordgo.Session, in *discordgo.InteractionCreate){
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData {
			Content: "Interaction received",
		},
	}
	
	err := session.InteractionRespond(in.Interaction, &data)
	if err != nil {
		log.Panicf("Error while interaction respond, %v", err)
	}

	ch, err := session.MessageThreadStart(s.ChannelID, in.Message.ID, "new-thread", 0)
	if err != nil {
		log.Panicf("Error while starting thread, %v", err)
	}

	log.Printf("Successfully created thread id: %v", ch.ID)
	
	time.Sleep(time.Second)
	err = session.InteractionResponseDelete(in.Interaction)
	if err != nil {
		log.Panicf("Error while interaciton response delete, %v", err)
	}
	
}

func (s *Service) RespondMessage(session *discordgo.Session, msg *discordgo.MessageCreate){
	if msg.ChannelID == s.ChannelID && !msg.Author.Bot { 
		_, err := session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v says: %v", msg.Author.Username, msg.Content))
		if err != nil{
			log.Panicf("Error while channel message send, %v", err)
		}
	}
	if msg.Content == "end" && msg.ChannelID != s.ChannelID { 
		_, err := session.ChannelDelete(msg.ChannelID)
		if err != nil{
			log.Panicf("Error while channel delete, %v", err)
		}
	}
}
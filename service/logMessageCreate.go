package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) LogMessageCreate(session *discordgo.Session, msg *discordgo.MessageCreate){
	if !msg.Author.Bot { 
		_, err := session.ChannelMessageSend(msg.ChannelID, fmt.Sprintf("%v says: `%v`", msg.Author.Username, msg.Content))
		if err != nil{
			log.Panicf("Error while channel message send, %v", err)
		}
	}
}
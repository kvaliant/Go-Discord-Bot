package service

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) LogReactionAdd(session *discordgo.Session, e *discordgo.MessageReactionAdd){
	msg, err := session.ChannelMessage(e.ChannelID, e.MessageID)
	if err != nil{
		log.Panicf("Error while fetching channel message, %v", err)
	}

	_, err = session.ChannelMessageSend(e.ChannelID, fmt.Sprintf("%v: `%v` \n %v reacted: %v",msg.Author.Username, msg.Content, e.Member.User.Username, e.Emoji.Name))
	if err != nil{
		log.Panicf("Error while channel message send, %v", err)
	}
}
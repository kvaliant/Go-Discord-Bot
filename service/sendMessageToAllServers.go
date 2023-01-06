package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type Guild struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func FetchGuildsList (session *discordgo.Session) ([]Guild, error) {
	req, err := http.NewRequest("GET", "https://discordapp.com/api/users/@me/guilds", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", session.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var guilds []Guild
	err = json.NewDecoder(resp.Body).Decode(&guilds)
	if err != nil {
		return nil, err
	}

	return guilds, nil
}

func (s *Service) SendMessageToAllServers(session *discordgo.Session, content string){
	guilds, err := FetchGuildsList(session)
	if err != nil {
		log.Panicf("Error while fetching guilds list, %v", err)
	}

	for _, element := range guilds {
		guildChannels, err := session.GuildChannels(element.Id)
		if err != nil {
			log.Panicf("Error while channel message send complex, %v", err)
		}
		
		for _, channel := range guildChannels {
			if channel.Type == discordgo.ChannelTypeGuildText {
				_, err := session.ChannelMessageSend(channel.ID, content)
				if err != nil {
					log.Panicf("Error while channel message send complex, %v", err)
				}			
			}
		}
	}
}
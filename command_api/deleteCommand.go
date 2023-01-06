package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DeleteCommand(){
	commandID := "1057635310059671573"
	applicationID := os.Getenv("APPLICATION")
	uri := fmt.Sprintf("https://discord.com/api/v10/applications/%v/commands/%v", applicationID, commandID)
	req, err := http.NewRequest("DELETE", uri, nil)
	if err != nil {
		log.Fatalf("Err while constructing http request, %v", err)
	}

	req.Header.Set("Accept", "application/json")
	tokenID := os.Getenv("TOKEN")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %v",tokenID))

	log.Printf("Sending http request, token : %v , applicationID : %v", tokenID, applicationID)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Err sending http request, %v", err)
	}

	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Err while reading response body, %v", err)
	}	
	
	log.Printf("resp : %v", string(res))
}
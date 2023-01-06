package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func PostCommand(){
	values := map[string]interface{}{
		"name": "blep",
		"type": "1",
		"description": "Send a random adorable animal photo",
	}
	byte := new(bytes.Buffer)
	json.NewEncoder(byte).Encode(values)

	applicationID := os.Getenv("APPLICATION")
	uri := fmt.Sprintf("https://discord.com/api/v10/applications/%v/commands", applicationID)
	req, err := http.NewRequest("POST", uri, byte)
	if err != nil {
		log.Fatalf("Err while constructing http request, %v", err)
	}

	tokenID := os.Getenv("TOKEN")
	req.Header.Set("Authorization", fmt.Sprintf("Bot %v",tokenID))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

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

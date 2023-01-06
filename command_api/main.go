package main

import "github.com/kvaliant/go-discord-bot/config"

func init() {
	config.LoadEnvVariables()
}

func main() {
	GetCommand()
	// PostCommand()
	// DeleteCommand()
}
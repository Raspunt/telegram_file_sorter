package main

import (
	"log"
	bot "telega/src/Telegram_Bot"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	bot.BotRun()

}

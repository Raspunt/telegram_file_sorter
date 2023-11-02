package telegrambot

import (
	"fmt"
	"log"
	"os"
	"telega/src/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func BotRun() {

	BotToken := os.Getenv("BotToken")

	fmt.Println(BotToken)

	bot, err := tgbotapi.NewBotAPI(BotToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			doc := update.Message.Document

			if doc != nil{

				fileID := doc.FileID

				file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
				if err != nil {
                    log.Println("Failed to get file:", err)
                    continue
                }

				fmt.Println(12,file,12)
				
			}

			switch msg.Text {

			case "create_keyboard":
				CreateKeyboard(bot, msg)
			case "lessons":

				lessons := db.ListOfLessons()

				for _, lesson := range lessons {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, lesson.Name)
					bot.Send(msg)
				}

			case "add_file_to_lesson":

			}

		}
	}

}

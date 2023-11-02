package telegrambot

import (
	"fmt"
	"log"
	db "telega/src/db"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var commands_row = tgbotapi.NewKeyboardButtonRow(
	tgbotapi.NewKeyboardButton("create_keyboard"),
	tgbotapi.NewKeyboardButton("lessons"),
	tgbotapi.NewKeyboardButton("add_file_to_lesson"),
)

func GetKeyboard() tgbotapi.ReplyKeyboardMarkup {

	listofLessons := CreateListLessons()

	fmt.Println(listofLessons)

	var numericKeyboard = tgbotapi.NewReplyKeyboard(
		commands_row,
		listofLessons,
	)

	return numericKeyboard

}

func CreateListLessons() []tgbotapi.KeyboardButton {

	var btnLists []tgbotapi.KeyboardButton
	// var rows []tgbotapi.KeyboardButton

	for _, lesson := range db.ListOfLessons() {
		fmt.Println(lesson.Name)

		btnLists = append(btnLists, tgbotapi.NewKeyboardButton(lesson.Name))
	}

	keyboard := tgbotapi.NewKeyboardButtonRow(btnLists...)
	return keyboard
}

func CreateKeyboard(bot *tgbotapi.BotAPI, msg tgbotapi.MessageConfig) {

	msg.ReplyMarkup = GetKeyboard()

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

}

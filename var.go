package main

import botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var (
	// Endpoint api endpoint telegram is all interface prefixes
	endpoint = "https://api.telegram.org/bot%s/%s"
	// Bot bot api
	bot *botapi.BotAPI
	// Menu texts
	firstMenu  = "<b>Welcome</b>\n\nI'm go_hello_bot, can I help you."
	secondMenu = "<b>Help Menu</b>\n\nIf you have any questions, please contact the owner github!!."

	// Button texts
	nextButton     = "Help"
	backButton     = "Back"
	tutorialButton = "Github"

	// Keyboard layout for the first menu. One button, one row
	firstMenuMarkup = botapi.NewInlineKeyboardMarkup(
		botapi.NewInlineKeyboardRow(
			botapi.NewInlineKeyboardButtonData(nextButton, nextButton),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = botapi.NewInlineKeyboardMarkup(
		botapi.NewInlineKeyboardRow(
			botapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
		botapi.NewInlineKeyboardRow(
			botapi.NewInlineKeyboardButtonURL(tutorialButton, "https://github.com/EscAlice"),
		),
	)
)

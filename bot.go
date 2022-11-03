package main

import (
	"context"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"golang.org/x/net/proxy"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Socks5Client(socks5, token string) (*botapi.BotAPI, error) {
	client := &http.Client{}
	tgProxyURL, err := url.Parse(socks5)
	if err != nil {
		log.Printf("Failed to parse proxy URL:%s\n", err)
	}
	tgDialer, err := proxy.FromURL(tgProxyURL, proxy.Direct)
	if err != nil {
		log.Printf("Failed to obtain proxy dialer: %s\n", err)
	}
	tgTransport := &http.Transport{
		Dial: tgDialer.Dial,
	}
	client.Transport = tgTransport
	bot, err = botapi.NewBotAPIWithClient(token, endpoint, client)
	if err != nil {
		log.Panic(err)
	}
	return bot, nil
}

func receiveUpdates(ctx context.Context, updates botapi.UpdatesChannel) {
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update botapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break
	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *botapi.Message) {
	user := message.From
	text := message.Text
	if user == nil {
		return
	}
	// Print to console
	log.Printf("%s wrote %s", user.FirstName, text)
	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else {
		copyMsg := botapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}
	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

// get a command, react accordingly
func handleCommand(chatId int64, command string) error {
	var err error
	switch command {
	case "/menu":
		err = sendMenu(chatId)
		break
	}
	return err
}

func handleButton(query *botapi.CallbackQuery) {
	var text string

	markup := botapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == nextButton {
		text = secondMenu
		markup = secondMenuMarkup
	} else if query.Data == backButton {
		text = firstMenu
		markup = firstMenuMarkup
	}
	callbackCfg := botapi.NewCallback(query.ID, "")
	_, _ = bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := botapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = botapi.ModeHTML
	_, _ = bot.Send(msg)
}

func sendMenu(chatId int64) error {
	msg := botapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = botapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}

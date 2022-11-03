package main

import (
	"bufio"
	"context"
	"flag"
	botapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var (
	socks5 string
	token  string
	help   bool
)

func init() {
	flag.StringVar(&token, "token", "", "a token is a string that authenticates your bot (not your account) on the bot API.")
	flag.StringVar(&socks5, "socks5", "", "the proxy.")
	flag.BoolVar(&help, "help", false, "help")
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	var err error
	if token == "" {
		log.Print("token is nil")
		return
	}
	switch socks5 {
	case "":
		bot, err = botapi.NewBotAPI(token)
		if err != nil {
			log.Panic(err)
			return
		}
	default:
		bot, err = Socks5Client(socks5, token)
		if err != nil {
			log.Panic(err)
			return
		}
	}
	// Set this to true to log all interactions with telegram servers
	bot.Debug = true
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := botapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	go receiveUpdates(ctx, updates)
	log.Println("Start listening for updates. Press enter to stop")
	// Wait for a newline symbol, then cancel handling updates
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()
}

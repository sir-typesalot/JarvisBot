package main

import (
	"jarvis-bot/bot"
)

func main() {
	bot.Run()
	<-make(chan struct{})
	return
}

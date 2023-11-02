package main

import (
	"fmt"

	"github.com/DiegoAraujoJS/health-bot/environment"
	"github.com/DiegoAraujoJS/health-bot/health"
	"github.com/DiegoAraujoJS/health-bot/messages"
)

func main() {
    env := environment.GetSafeEnvVariables()

    health.TickHealthChecker([]func(message *messages.AlarmMessage){
        messages.SendTelegramMessageToPhoneId(env.TelegramBotToken, env.PhoneId),
        messages.SendDiscordMessage(env.DiscordBotToken, env.DiscordChannelIds),
    })
    fmt.Println("Started to tick health checker")

    <- make(chan bool)
}

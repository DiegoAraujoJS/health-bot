package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DiegoAraujoJS/health-bot/environment"
)

func SendTelegramMessageToPhoneId(phoneId string) func(messsage *AlarmMessage) {
    return func(message *AlarmMessage) {
        sendTelegramMessage(message, phoneId)
    }
}

func sendTelegramMessage(message *AlarmMessage, to string) {
    env := environment.GetSafeEnvVariables()
    botToken := env.TelegramBotToken

    body, _ := json.Marshal(map[string]string {
        "chat_id": to,
        "text": message.Title + "\n\n" + message.Description + "\n\n" + message.EasyTest,
    })
    _, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), "application/json", bytes.NewBuffer(body))
    if err != nil {
        return
    }
}

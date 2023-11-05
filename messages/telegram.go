package messages

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func SendTelegramMessageToPhoneId(botToken string, phoneId string) func(messsage *AlarmMessage) {
    return func(message *AlarmMessage) {
        sendTelegramMessage(botToken, message, phoneId)
    }
}

func sendTelegramMessage(botToken string, message *AlarmMessage, to string) {
    body, _ := json.Marshal(map[string]string {
        "chat_id": to,
        "text": message.Title + "\n\n" + message.Description + "\n\n" + message.EasyTest,
    })
    _, err := HttpClient.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), "application/json", bytes.NewBuffer(body))
    if err != nil {
        return
    }
}

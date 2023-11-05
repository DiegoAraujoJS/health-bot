package messages

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func SendDiscordMessage(discordBotToken string, discordChannelIds []string) func (message *AlarmMessage) {
    return func(message *AlarmMessage) {
        sendDiscordMessage(discordBotToken, discordChannelIds, message)
    }
}

func sendDiscordMessage(discordBotToken string, discordChannelIds []string, message *AlarmMessage) {
    dg, err := discordgo.New("Bot " + discordBotToken)
    if err != nil {
        fmt.Println("Error creating Discord session,", err)
        return
    }

    // Send a message to the specified channels
    for _, discordChannelId := range discordChannelIds {
        _, err = dg.ChannelMessageSendEmbed(discordChannelId, &discordgo.MessageEmbed{
            Title: message.Title,
            Description: message.Description + "\n\n" + message.EasyTest,
            Color: 2123412,
        })
        if err != nil {
            fmt.Println("Error sending message,", err)
            continue
        }
    }

    // Close the Discord session
    dg.Close()
}

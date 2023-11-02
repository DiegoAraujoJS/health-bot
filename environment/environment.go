package environment

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type Configuration struct {
    PhoneId             string  `json:"PHONE_ID"`
    TelegramBotToken    string  `json:"TELEGRAM_BOT_TOKEN"`
    PingUrl             string  `json:"PING_URL"`
    DiscordBotToken     string  `json:"DISCORD_BOT_TOKEN"`
    DiscordChannelIds   []string  `json:"DISCORD_CHANNEL_IDS"`
}

var Env *Configuration

func GetSafeEnvVariables() *Configuration {
    if Env != nil {return Env}

    body, err := os.ReadFile("./.env.json")
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(body, &Env)
    if err != nil {
        panic(err)
    }

    v := reflect.ValueOf(*Env)

    for i := 0; i < v.NumField(); i++ {
        if v.Field(i).IsZero() {
            panic(fmt.Sprintf("empty .env.json configuration value: %v", v.Type().Field(i).Tag.Get("json")))
        }
    }

    return Env
}

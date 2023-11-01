package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    ticker := time.NewTicker(60 * time.Second)

    var isInErrorState = false

    go func() {
        for t := range ticker.C {

            if isInErrorState {
                pingResult := GetGitDeployHealth()
                if pingResult {
                    isInErrorState = false
                }
                continue
            }

            pingResult := GetGitDeployHealth()
            if !pingResult {
                isInErrorState = true
                sendMessage(fmt.Sprintf("Request to %v failed! at %v", os.Getenv("PING_URL"), t.String()), os.Getenv("PHONE_ID"))
            }

        }
    }()

    <- make(chan bool)
}

func sendMessage(message string, to string) {
    botToken := os.Getenv("BOT_TOKEN")
    body, _ := json.Marshal(map[string]string {
        "chat_id": to,
        "text": message,
    })
    fmt.Println(string(body))
    _, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken), "application/json", bytes.NewBuffer(body))
    if err != nil {
        return
    }
}

func GetGitDeployHealth() bool {
    pingUrl := os.Getenv("PING_URL")
    resp, err := http.Get(pingUrl)
    if err != nil {
        return false
    }
    return resp.StatusCode == 200
}

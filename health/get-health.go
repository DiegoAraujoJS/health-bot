package health

import (
	"github.com/DiegoAraujoJS/health-bot/messages"
)

func getUrlHealth(pingUrl string) bool {
    resp, err := messages.HttpClient.Get(pingUrl)
    if err != nil {
        return false
    }
    return resp.StatusCode == 200
}

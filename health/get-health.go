package health

import (
	"net/http"

	"github.com/DiegoAraujoJS/health-bot/environment"
)

func getGitDeployHealth() bool {
    env := environment.GetSafeEnvVariables()
    pingUrl := env.PingUrl

    resp, err := http.Get(pingUrl)
    if err != nil {
        return false
    }
    return resp.StatusCode == 200
}

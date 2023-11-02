package health

import (
	"fmt"
	"time"

	"github.com/DiegoAraujoJS/health-bot/environment"
	"github.com/DiegoAraujoJS/health-bot/messages"
)


func TickHealthChecker(onUnhealthyResponse []func (message *messages.AlarmMessage)) {
    env := environment.GetSafeEnvVariables()

    ticker := time.NewTicker(60 * time.Second)

    var isInErrorState = false

    go func() {

        defer func() {
            if r := recover(); r != nil {
                fmt.Println(r)
            }
        }()

        for t := range ticker.C {

            if isInErrorState {
                pingResult := getGitDeployHealth()
                if pingResult {
                    isInErrorState = false
                }
                continue
            }

            pingResult := getGitDeployHealth()
            if !pingResult {
                isInErrorState = true
                for _, f := range onUnhealthyResponse {
                    f(&messages.AlarmMessage{
                        Title: "Error en Git Deploy",
                        Description: "El servidor no responde. Probablemente esté caído o se haya reseteado la máquina de test",
                        EasyTest: env.PingUrl,
                        Time: t,
                    })
                }
            }

        }
    }()
}

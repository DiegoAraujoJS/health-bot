package health

import (
	"fmt"
	"log"
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

            pingResult := getUrlHealth(env.PingUrl)
            if isInErrorState {
                if pingResult {
                    isInErrorState = false
                }
                continue
            }

            if !pingResult {
                isInErrorState = true
                for _, f := range onUnhealthyResponse {
                    func () {

                        defer func() {
                            if r := recover(); r != nil {
                                log.Println("error while sending message:", r)
                            }
                        }()

                        f(&messages.AlarmMessage{
                            Title: "Error en Git Deploy",
                            Description: "El servidor no responde. Probablemente esté caído o se haya reseteado la máquina de test",
                            EasyTest: env.PingUrl,
                            Time: t,
                        })
                    }()
                }
            }
        }
    }()
}

package messages

import "time"

type AlarmMessage struct {
    Title       string
    Description string
    EasyTest    string
    AppUrl      string
    Time        time.Time
}

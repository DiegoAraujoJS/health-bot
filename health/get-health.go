package health

import (
	"net/http"
)

func getUrlHealth(pingUrl string) bool {
    resp, err := http.Get(pingUrl)
    if err != nil {
        return false
    }
    return resp.StatusCode == 200
}

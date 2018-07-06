package tool

import (
	"net/http"
	"time"
	"errors"
)

func PingServer() error {
	for i := 0; i < 10; i++ {
		resp, err := http.Get("http://localhost:8080/monitor/ping")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("server error")
}
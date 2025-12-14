package checker

import (
	"net/http"
	"time"
)

func CheckURL(url string) string {
	if len(url) > 0 && url[:4] != "http" {
		url = "http://" + url
	}

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return "not available"
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return "available"
	}
	return "not available"
}

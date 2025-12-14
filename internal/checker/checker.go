package checker

import (
	"net/http"
	"time"
)

var defaultTimeout = 5 * time.Second

func SetDefaultTimeout(timeout time.Duration) {
	defaultTimeout = timeout
}

func CheckURL(url string) string {
	if len(url) > 0 && url[:4] != "http" {
		url = "http://" + url
	}

	client := &http.Client{Timeout: defaultTimeout}

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

package utils

import (
	"crypto/tls"
	"github.com/BaizeSec/Pusher/common"
	"net/http"
	"strings"
	"time"
)

func PostJson(url string, body string) bool {
	body = conv(body)

	client := http.Client{
		Timeout: time.Duration(common.TIMEOUT) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(body))

	if err != nil {
		return false
	}

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	resp, err := client.Do(req)

	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return true
}

func conv(body string) string {
	body = strings.ReplaceAll(body, "\n", "\\n")
	return body
}

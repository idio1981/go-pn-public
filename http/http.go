package http

import (
	"io"
	"net/http"
	"time"

	"github.com/idio1981/go-pn-public/logger"
)

func Get(url string) (int, string, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Error making GET request: %s", err)
		return 0, "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body: %s", err)
		return 0, "", err
	}

	return resp.StatusCode, string(body), nil
}

func Post(url string, body io.Reader) (int, string, error) {
	client := &http.Client{
		Timeout: 2 * time.Minute,
	}

	// Make the HTTP POST request
	resp, err := client.Post(url, "application/json", body)
	if err != nil {
		logger.Error("Error making POST request:", err)
		return 0, "", err
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Error reading response body:", err)
		return 0, "", err
	}

	return resp.StatusCode, string(respBody), nil
}

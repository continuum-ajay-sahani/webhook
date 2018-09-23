package util

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ContinuumLLC/platform-correlation-engine/src/service"
)

// RequestData contain request specific data
type RequestData struct {
}

// Call generic method to call rest api
func Call(method, url string, a *Response) (*http.Response, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := service.NewClient()

	return client.Do(req)
}

// Response represent res data
type Response struct {
	RID          string
	RandDuration time.Duration
}

package exporter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var DefaultClient = &http.Client{Timeout: 1 * time.Minute}

func HTTPRequest[T any](endpoint string, path string, headers map[string]string, formData url.Values) (T, error) {
	var output T

	req, err := http.NewRequest("POST", endpoint+path, strings.NewReader(formData.Encode()))
	if err != nil {
		return output, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := DefaultClient.Do(req)
	if err != nil {
		return output, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return output, fmt.Errorf("status: %s, body: %s", resp.Status, string(body))
	}

	if err := json.Unmarshal(body, &output); err != nil {
		return output, fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	return output, nil
}

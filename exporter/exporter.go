package exporter

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type Metrics struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric map[string]string `json:"metric"`
			Value  []any             `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type LabelValues struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}

func CheckIfMetricsExists(endpoint string, headers map[string]string, metricsList []string) ([]string, error) {
	labelValues, err := HTTPRequest[LabelValues](endpoint, "/api/v1/label/__name__/values", headers, url.Values{})
	if err != nil {
		return nil, err
	}

	// Create a map of available metrics for O(1) lookup
	availableMetrics := make(map[string]bool)
	for _, metric := range labelValues.Data {
		availableMetrics[metric] = true
	}

	// Check which requested metrics exist
	var existingMetrics []string
	for _, metric := range metricsList {
		if availableMetrics[metric] {
			existingMetrics = append(existingMetrics, metric)
		}
	}

	return existingMetrics, nil
}

func GetMetrics(endpoint string, headers map[string]string, metricsList []string, queryTime *time.Time) (*Metrics, error) {
	formData := url.Values{}

	if len(metricsList) == 0 {
		// If no metrics specified, query all metrics
		formData.Set("query", `{__name__!=""}`)
	} else {
		// Build query for specific metrics
		query := fmt.Sprintf(`{__name__=~"%s"}`, strings.Join(metricsList, "|"))
		formData.Set("query", query)
	}

	// Add time parameter if provided
	if queryTime != nil {
		formData.Set("time", fmt.Sprintf("%d", queryTime.Unix()))
	}

	return HTTPRequest[*Metrics](endpoint, "/api/v1/query", headers, formData)
}

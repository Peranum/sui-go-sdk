package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// QueryEvents sends a query to the specified URL to retrieve event data
// for the given sender address, cursor, limit, and order preference.
// It returns the raw JSON response as a map.
func QueryEventsBySender(url, senderAddress string, cursor string, limit int, descendingOrder bool) (map[string]interface{}, error) {

	// Prepare the request body as a map to convert to JSON
	requestBody := map[string]interface{}{
		"query": map[string]interface{}{
			"Sender": senderAddress,
		},
		"cursor":           cursor,
		"limit":            limit,
		"descending_order": descendingOrder,
	}

	// Marshal the request body into JSON format
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %w", err)
	}

	// Send the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the response body into a map
	var jsonResponse map[string]interface{}
	if err := json.Unmarshal(body, &jsonResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return jsonResponse, nil
}

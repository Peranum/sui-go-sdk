package staking

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetStakesByIds fetches one or more delegated stakes for the given staked Sui IDs from the specified Sui node URL.
// If a stake was withdrawn, its status will be Unstaked.
func GetStakesByIds(url string, stakedSuiIDs []string) (map[string]interface{}, error) {

	// Prepare the request body as a map to convert to JSON
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "suix_getStakesByIds",
		"params":  []interface{}{stakedSuiIDs},
		"id":      1,
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

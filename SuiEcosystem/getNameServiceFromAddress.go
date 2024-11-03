package suiecosystem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ResolveNameServiceNames performs an HTTP request to retrieve resolved names for a Sui address
func ResolveNameServiceNames(url, address string, cursor string, limit uint) (map[string]interface{}, error) {
	// Prepare the request payload with the RPC method and parameters
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_resolveNameServiceNames",
		"params":  []interface{}{address},
	}

	// Add optional parameters to the request payload
	if cursor != "" {
		requestBody["params"] = append(requestBody["params"].([]interface{}), cursor)
	} else {
		requestBody["params"] = append(requestBody["params"].([]interface{}), nil)
	}

	requestBody["params"] = append(requestBody["params"].([]interface{}), limit)

	// Marshal the request payload to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request body: %w", err)
	}

	// Perform the HTTP POST request
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

	// Unmarshal the JSON response into a generic map
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return result, nil
}

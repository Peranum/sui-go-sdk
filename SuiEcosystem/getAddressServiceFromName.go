package suiecosystem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ResolveNameServiceAddress performs an HTTP request to resolve a name to a Sui address
func ResolveNameServiceAddress(url, name string) (map[string]interface{}, error) {
	// Prepare the request payload with the RPC method and parameters
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "suix_resolveNameServiceAddress",
		"params":  []interface{}{name},
	}

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

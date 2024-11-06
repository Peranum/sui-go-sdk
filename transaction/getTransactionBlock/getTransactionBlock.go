package transaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/Peranum/sui-go-sdk/shared"
)

// SuiTransactionRequestOptions represents options for the transaction request.
type SuiTransactionRequestOptions struct {
	ShowInput         bool `json:"showInput"`
	ShowRawInput      bool `json:"showRawInput"`
	ShowEffects       bool `json:"showEffects"`
	ShowEvents        bool `json:"showEvents"`
	ShowObjectChanges bool `json:"showObjectChanges"`
	ShowBalanceChanges bool `json:"showBalanceChanges"`
	ShowRawEffects    bool `json:"showRawEffects"`
}

// SuiTransactionResponse represents the structure of the response from the API.
type SuiTransactionResponse struct {
	Jsonrpc string                 `json:"jsonrpc"`
	ID      int                    `json:"id"`
	Result  map[string]interface{} `json:"result"`
	Error   interface{}            `json:"error"`
}

// getTransactionBlock sends a JSON-RPC request to fetch the transaction block details.
func getTransactionBlock(digest string, options SuiTransactionRequestOptions) (*SuiTransactionResponse, error) {
	// Construct JSON-RPC request body
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "sui_getTransactionBlock",
		"params":  []interface{}{digest, options},
	}

	// Encode the request body to JSON
	requestBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error encoding request: %v", err)
	}

	// Send POST request
	resp, err := http.Post(shared.SUI_NODE_URL, "application/json", bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Read and decode the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var transactionResponse SuiTransactionResponse
	if err := json.Unmarshal(body, &transactionResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// Check if an error was returned in the response
	if transactionResponse.Error != nil {
		return nil, fmt.Errorf("error in response: %v", transactionResponse.Error)
	}

	return &transactionResponse, nil
}

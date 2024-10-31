package coins

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "sui-go-sdk/shared"
)

// CoinMetadata represents the metadata structure for a coin.
type CoinMetadata struct {
    Decimals    uint8   `json:"decimals"`
    Name        string  `json:"name"`
    Symbol      string  `json:"symbol"`
    Description string  `json:"description"`
    IconURL     *string `json:"iconUrl"` // Optional field, can be null
    ID          *string `json:"id"`      // Optional field, can be null
}

// GetCoinMetadata retrieves the metadata for a specific coin type.
func GetCoinMetadata(url, coinType string) (CoinMetadata, error) {
    // Create the request payload
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getCoinMetadata",
        Params:  []interface{}{coinType},
    }

    // Convert the request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return CoinMetadata{}, err
    }

    // Make the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return CoinMetadata{}, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return CoinMetadata{}, err
    }

    // Parse the JSON response into shared.RPCResponse
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return CoinMetadata{}, err
    }

    // Attempt to convert rpcResponse.Result into CoinMetadata
    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return CoinMetadata{}, err
    }

    var metadata CoinMetadata
    if err := json.Unmarshal(resultBytes, &metadata); err != nil {
        return CoinMetadata{}, err
    }

    return metadata, nil
}

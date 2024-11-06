package balance

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "github.com/Peranum/sui-go-sdk/shared"
)

// TotalSupply represents the total supply structure for a coin.
type TotalSupply struct {
    Value string `json:"value"` // BigInt is represented as a string in JSON
}

// GetTotalSupply retrieves the total supply for a specific coin type.
func GetTotalSupply(url, coinType string) (TotalSupply, error) {
    // Create the request payload
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getTotalSupply",
        Params:  []interface{}{coinType},
    }

    // Convert the request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return TotalSupply{}, err
    }

    // Make the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return TotalSupply{}, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return TotalSupply{}, err
    }

    // Parse the JSON response into shared.RPCResponse
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return TotalSupply{}, err
    }

    // Attempt to convert rpcResponse.Result into TotalSupply
    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return TotalSupply{}, err
    }

    var supply TotalSupply
    if err := json.Unmarshal(resultBytes, &supply); err != nil {
        return TotalSupply{}, err
    }

    return supply, nil
}

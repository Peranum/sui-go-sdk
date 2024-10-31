package balance

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "sui-go-sdk/shared"
)

// Define the structure for the balance result
type CoinBalance struct {
    CoinType        string            `json:"coinType"`
    CoinObjectCount int               `json:"coinObjectCount"`
    TotalBalance    string            `json:"totalBalance"`
    LockedBalance   map[string]string `json:"lockedBalance"`
}

// GetBalance retrieves the balance for a specific coin type for the provided address.
func GetBalance(url, owner, coinType string) (CoinBalance, error) {
    // Prepare parameters with optional coin type
    params := []interface{}{owner}
    if coinType != "" {
        params = append(params, coinType)
    }

    // Create the request payload
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getBalance",
        Params:  params,
    }

    // Convert the request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return CoinBalance{}, err
    }

    // Make the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return CoinBalance{}, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return CoinBalance{}, err
    }

    // Parse the JSON response
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return CoinBalance{}, err
    }

    // Attempt to convert rpcResponse.Result into CoinBalance
    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return CoinBalance{}, err
    }

    var balance CoinBalance
    if err := json.Unmarshal(resultBytes, &balance); err != nil {
        return CoinBalance{}, err
    }

    return balance, nil
}

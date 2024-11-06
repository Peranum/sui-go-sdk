package coins

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "github.com/Peranum/sui-go-sdk/shared"
)

type Coin struct {
    CoinType           string `json:"coinType"`
    CoinObjectID       string `json:"coinObjectId"`
    Version            string `json:"version"`
    Digest             string `json:"digest"`
    Balance            string `json:"balance"`
    PreviousTransaction string `json:"previousTransaction"`
}

type CoinPage struct {
    Data         []Coin  `json:"data"`
    NextCursor   *string `json:"nextCursor"`  // Pointer to handle null value
    HasNextPage  bool    `json:"hasNextPage"`
}

// GetAllCoins retrieves all coins owned by an address with optional pagination.
func GetAllCoins(url, owner string, cursor *string, limit *int) (CoinPage, error) {
    // Prepare parameters with optional cursor and limit
    params := []interface{}{owner}
    if cursor != nil {
        params = append(params, *cursor)
    } else {
        params = append(params, nil) // For JSON-RPC, this represents a null value
    }
    if limit != nil {
        params = append(params, *limit)
    }

    // Create the request payload
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getAllCoins",
        Params:  params,
    }

    // Convert the request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return CoinPage{}, err
    }

    // Make the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return CoinPage{}, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return CoinPage{}, err
    }

    // Parse the JSON response into shared.RPCResponse
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return CoinPage{}, err
    }

    // Attempt to convert rpcResponse.Result into CoinPage
    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return CoinPage{}, err
    }

    var coinPage CoinPage
    if err := json.Unmarshal(resultBytes, &coinPage); err != nil {
        return CoinPage{}, err
    }

    return coinPage, nil
}

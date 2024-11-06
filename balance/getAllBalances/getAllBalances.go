package balance

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "github.com/Peranum/sui-go-sdk/shared" // Import shared package
)

type Balance struct {
    CoinType        string            `json:"coinType"`
    CoinObjectCount int               `json:"coinObjectCount"`
    TotalBalance    string            `json:"totalBalance"`
    LockedBalance   map[string]string `json:"lockedBalance"`
}

//Return the total coin balance for all coin type, owned by the address owner.
func GetAllBalances(url string, owner string) ([]Balance, error) {

    // Create request payload using shared.RPCRequest
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getAllBalances",
        Params:  []interface{}{owner},
    }

    // Convert the request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return nil, err
    }

    // Make the HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    // Parse the JSON response into shared.RPCResponse
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return nil, err
    }

    // Type assertion: ensure rpcResponse.Result is a slice of Balance
    resultList, ok := rpcResponse.Result.([]interface{})
    if !ok {
        return nil, err
    }

    // Convert resultList items to []Balance
    var balances []Balance
    for _, item := range resultList {
        itemBytes, err := json.Marshal(item)
        if err != nil {
            return nil, err
        }

        var balance Balance
        if err := json.Unmarshal(itemBytes, &balance); err != nil {
            return nil, err
        }
        balances = append(balances, balance)
    }

    return balances, nil
}

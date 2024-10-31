package balance

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "sui-go-sdk/shared" // Import shared package
    "fmt"
)

// Define the structure for the balance result
type Balance struct {
    CoinType        string            `json:"coinType"`
    CoinObjectCount int               `json:"coinObjectCount"`
    TotalBalance    string            `json:"totalBalance"`
    LockedBalance   map[string]string `json:"lockedBalance"`
}

// GetAllBalancesAndDetails first retrieves all balances and then fetches detailed information for each token type.
func GetAllBalancesAndDetails(url string, owner string) ([]Balance, error) {
    // Retrieve all balances
    allBalances, err := getAllBalances(url, owner)
    if err != nil {
        return nil, err
    }

    // Fetch additional details for each coin type
    var detailedBalances []Balance
    for _, balance := range allBalances {
        detailedBalance, err := getBalance(url, owner, balance.CoinType)
        if err != nil {
            fmt.Printf("Error fetching details for coin type %s: %v\n", balance.CoinType, err)
            continue
        }
        detailedBalances = append(detailedBalances, detailedBalance)
    }

    return detailedBalances, nil
}

// getAllBalances returns the total balance for all coin types owned by the address owner.
func getAllBalances(url string, owner string) ([]Balance, error) {
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getAllBalances",
        Params:  []interface{}{owner},
    }

    requestBody, err := json.Marshal(request)
    if err != nil {
        return nil, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return nil, err
    }

    resultList, ok := rpcResponse.Result.([]interface{})
    if !ok {
        return nil, fmt.Errorf("unexpected result format")
    }

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

// getBalance retrieves detailed balance information for a specific coin type for the provided address.
func getBalance(url, owner, coinType string) (Balance, error) {
    params := []interface{}{owner}
    if coinType != "" {
        params = append(params, coinType)
    }

    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getBalance",
        Params:  params,
    }

    requestBody, err := json.Marshal(request)
    if err != nil {
        return Balance{}, err
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return Balance{}, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Balance{}, err
    }

    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return Balance{}, err
    }

    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return Balance{}, err
    }

    var balance Balance
    if err := json.Unmarshal(resultBytes, &balance); err != nil {
        return Balance{}, err
    }

    return balance, nil
}

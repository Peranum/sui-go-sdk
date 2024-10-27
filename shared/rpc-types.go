package shared

type RPCRequest struct {
    Jsonrpc string        `json:"jsonrpc"`
    ID      int           `json:"id"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
}

type RPCResponse struct {
    Jsonrpc string      `json:"jsonrpc"`
    Result  interface{} `json:"result"`
    ID      int         `json:"id"`
}

package shared

type RPCRequest struct {
    Jsonrpc string        `json:"jsonrpc"`
    ID      int           `json:"id"`
    Method  string        `json:"method"`
    Params  []interface{} `json:"params"`
}

type RPCError struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

type RPCResponse struct {
    Jsonrpc string      `json:"jsonrpc"`
    Result  interface{} `json:"result"`
    Error   *RPCError   `json:"error,omitempty"`
    ID      int         `json:"id"`
}


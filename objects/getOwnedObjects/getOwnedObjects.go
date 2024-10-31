package objects

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"
    "sui-go-sdk/shared"
)

// ObjectOwner represents the owner field in the SuiObjectResponse.
type ObjectOwner struct {
    AddressOwner string `json:"AddressOwner"`
}

// SuiObjectData represents the data structure of each owned object.
type SuiObjectData struct {
    ObjectID           string      `json:"objectId"`
    Version            string      `json:"version"`
    Digest             string      `json:"digest"`
    Type               string      `json:"type"`
    Owner              ObjectOwner `json:"owner"`
    PreviousTransaction string     `json:"previousTransaction"`
    StorageRebate      string      `json:"storageRebate"`
}

// SuiObjectResponse represents each object response in the result data.
type SuiObjectResponse struct {
    Data SuiObjectData `json:"data"`
}

// ObjectsPage represents the response structure for the owned objects query.
type ObjectsPage struct {
    Data        []SuiObjectResponse `json:"data"`
    HasNextPage bool                `json:"hasNextPage"`
    NextCursor  *string             `json:"nextCursor"`
}

// GetOwnedObjects retrieves the list of objects owned by a given address.
func GetOwnedObjects(url, address string, query map[string]interface{}, cursor *string, limit *int) (ObjectsPage, error) {
    // Define the query options to include type and owner
    queryOptions := map[string]interface{}{
        "showType":               true,
        "showOwner":              true,
        "showPreviousTransaction": true,
        "showStorageRebate":       true,
    }

    // Prepare the complete query object
    if query == nil {
        query = make(map[string]interface{})
    }
    query["options"] = queryOptions

    // Prepare parameters
    params := []interface{}{address, query}
    if cursor != nil {
        params = append(params, *cursor)
    } else {
        params = append(params, nil)
    }
    if limit != nil {
        params = append(params, *limit)
    }

    // Create JSON-RPC request
    request := shared.RPCRequest{
        Jsonrpc: "2.0",
        ID:      1,
        Method:  "suix_getOwnedObjects",
        Params:  params,
    }

    // Convert request to JSON
    requestBody, err := json.Marshal(request)
    if err != nil {
        return ObjectsPage{}, err
    }

    // Make HTTP request
    resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
    if err != nil {
        return ObjectsPage{}, err
    }
    defer resp.Body.Close()

    // Read the response
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return ObjectsPage{}, err
    }

    // Parse JSON response into shared.RPCResponse
    var rpcResponse shared.RPCResponse
    if err := json.Unmarshal(body, &rpcResponse); err != nil {
        return ObjectsPage{}, err
    }

    // Convert result to ObjectsPage
    resultBytes, err := json.Marshal(rpcResponse.Result)
    if err != nil {
        return ObjectsPage{}, err
    }

    var objectsPage ObjectsPage
    if err := json.Unmarshal(resultBytes, &objectsPage); err != nil {
        return ObjectsPage{}, err
    }

    return objectsPage, nil
}

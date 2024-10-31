package transaction

import (
	"context"
	"sui-go-sdk/shared"

	"github.com/machinebox/graphql"
)

// QueryTransactionBlockDetailsByDigest executes a GraphQL query to fetch TransactionBlock details by a given digest
func QueryTransactionBlockDetailsByDigest(digest string) (map[string]interface{}, error) {
	client := graphql.NewClient(shared.SuiGraphQLEndpoint) 

	query := `
		query ($digest: String!) {
			transactionBlock(digest: $digest) {
				gasInput {
					gasSponsor {
						address
					}
					gasPrice
					gasBudget
				}
				effects {
					status
					timestamp
					checkpoint {
						sequenceNumber
					}
					epoch {
						epochId
						referenceGasPrice
					}
				}
			}
		}
	`

	req := graphql.NewRequest(query)
	req.Var("digest", digest)

	// Context for the request
	ctx := context.Background()

	// The response is stored in map[string]interface{} format for flexible output
	var resp map[string]interface{}
	if err := client.Run(ctx, req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

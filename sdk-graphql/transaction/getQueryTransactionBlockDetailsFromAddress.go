package transaction

import (
	"context"
	"github.com/Peranum/sui-go-sdk/shared"
	"github.com/machinebox/graphql"
)
// QueryTransactionBlocksWithBalanceChanges performs a GraphQL query to fetch transaction blocks with balance changes
func QueryTransactionBlocksWithBalanceChanges(address string) (map[string]interface{}, error) {
	client := graphql.NewClient(shared.SuiGraphQLEndpoint)

	query := `
		query ($address: SuiAddress!) {
			transactionBlocks(filter: { affectedAddress: $address }) {
				nodes {
					digest
					effects {
						balanceChanges {
							nodes {
								owner {
									address
								}
								amount
							}
						}
					}
				}
			}
		}
	`

	req := graphql.NewRequest(query)
	req.Var("address", address)

	ctx := context.Background()
	var resp map[string]interface{}

	if err := client.Run(ctx, req, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}
package transaction

import (
	"context"
	"github.com/Peranum/sui-go-sdk/shared"
	"github.com/machinebox/graphql"
)

func QueryTransactionBlockDetails(address string) (map[string]interface{}, error) {
	client := graphql.NewClient(shared.SuiGraphQLEndpoint) 

	query := `
		query ($address: SuiAddress!) {
			address(address: $address) {
				transactionBlocks {
					nodes {
						digest
						sender {
							address
						}
						gasInput {
							objects {
								address
								owner
							}
							price
							budget
						}
						kind
						signatures
						effects {
							status
							created {
								address
							}
						}
						expiration
						bcs
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

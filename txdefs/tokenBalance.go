package txdefs

import (
	"encoding/json"
	"fmt"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Return the all books from an specific author
// GET method
var TokenBalance = tx.Transaction{
	Tag:         "TokenBalance",
	Label:       "Token Balance",
	Description: "Token Balance",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Tag:         "proprietario",
			Label:       "proprietario",
			Description: "proprietario",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		prop, _ := req["proprietario"].(string)

		// Prepare couchdb query
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType":   "Token",
				"proprietario": prop,
			},
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error searching for token", 500)
		}

		tokens := response.Result

		fmt.Printf("Response %+v", response)
		fmt.Println(".")
		quantidade := 0.0
		for i := 0; i < len(tokens); i++ {
			quantidade += tokens[i]["quantidade"].(float64)
			fmt.Printf("Token %s %f", tokens[i]["id"].(string), quantidade)
			//fmt.println(".")
		}

		balance := make(map[string]interface{})
		balance["balance"] = quantidade

		responseJson, err := json.Marshal(balance)

		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "error marshaling response", 500)
		}

		return responseJson, nil
	},
}

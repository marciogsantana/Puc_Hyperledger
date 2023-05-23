package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// POST Method
var CreateNewToken = tx.Transaction{
	Tag:         "createToken",
	Label:       "Create Token",
	Description: "Create a token",
	Method:      "POST",

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Id",
			Description: "Token Id",
			DataType:    "->Token",
			Required:    true,
		},

		{
			Tag:         "dono",
			Label:       "dono",
			Description: "owner token",
			DataType:    "->proprietario",
			Required:    true,
		},

		{
			Tag:         "quantidade",
			Label:       "quantidade",
			Description: "quantidade token",
			DataType:    "number",
			Required:    true,
		},

		{
			Tag:         "burn",
			Label:       "burn",
			Description: "burn token",
			DataType:    "boolean",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		id, _ := req["id"].(string)
		quantidade, _ := req["quantidade"].(float64)
		burn, _ := req["burn"].(bool)

		tokenMap := make(map[string]interface{})
		tokenMap["@assetType"] = "Token"
		tokenMap["@assetType"] = "proprietario"
		tokenMap["id"] = id
		tokenMap["quantidade"] = quantidade
		tokenMap["burn"] = burn

		tokenAsset, err := assets.NewAsset(tokenMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new token on channel
		_, err = tokenAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		tokenJSON, nerr := json.Marshal(tokenAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return tokenJSON, nil
	},
}

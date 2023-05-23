package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Library on channel
// POST Method
var CreateNewLibrary = tx.Transaction{
	Tag:         "createNewLibrary",
	Label:       "Create New Library",
	Description: "Create a New Library",
	Method:      "POST",
	Callers:     []string{"$org3MSP", "$orgMSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "name",
			Label:       "Name",
			Description: "Name of the library",
			DataType:    "string",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		name, _ := req["name"].(string)

		libraryMap := make(map[string]interface{})
		libraryMap["@assetType"] = "library"
		libraryMap["name"] = name

		libraryAsset, err := assets.NewAsset(libraryMap)
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new library on channel
		_, err = libraryAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		libraryJSON, nerr := json.Marshal(libraryAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return libraryJSON, nil
	},
}

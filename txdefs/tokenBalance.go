package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// GET method
var TokenBalance = tx.Transaction{
	Tag:         "BalanceToken",
	Label:       "BalanceToken",
	Description: "Total  Token",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Tag:         "proprietario",
			Label:       "Proprietário",
			Description: "Proprietário",
			DataType:    "->proprietario",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		proprietarioKey, ok := req["proprietario"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parametro proprietario deve ser um ativo.")
		}

		proprietarioAsset, errKey := proprietarioKey.Get(stub)
		if errKey != nil {
			return nil, errors.WrapError(errKey, "Falha ao obter ativo 'proprietário'.")
		}
		proprietarioMap := (map[string]interface{})(*proprietarioAsset)

		updatedProprietarioKey := make(map[string]interface{})
		updatedProprietarioKey["@assetType"] = "proprietario"
		updatedProprietarioKey["@key"] = proprietarioMap["@key"]

		// Prepara a consulta no CouchDB
		query := map[string]interface{}{
			"selector": map[string]interface{}{
				"@assetType":   "Token",
				"proprietario": updatedProprietarioKey,
			},
		}

		var err error
		response, err := assets.Search(stub, query, "", true)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "Erro ao buscar token por proprietário.", 500)
		}

		tokens := response.Result

		var quantidade float64 = 0

		for i := 0; i < len(tokens); i++ {
			if !tokens[i]["burned"].(bool) {
				quantidade = quantidade + tokens[i]["quantidade"].(float64)
			}
		}

		balance := make(map[string]interface{})
		balance["quantidade"] = quantidade

		responseJSON, err := json.Marshal(balance)
		if err != nil {
			return nil, errors.WrapErrorWithStatus(err, "Falha ao converter ativo para JSON.", 500)
		}

		return responseJSON, nil
	},
}

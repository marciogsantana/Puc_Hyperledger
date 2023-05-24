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
		proprietarioKey, ok := req["proprietario"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parametro proprietario deve ser um ativo.")
		}

		proprietarioAsset, err := proprietarioKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Falha ao obter ativo 'proprietário'.")
		}
		proprietarioMap := (map[string]interface{})(*proprietarioAsset)

		updatedProprietarioKey := make(map[string]interface{})
		updatedProprietarioKey["@assetType"] = "proprietario"
		updatedProprietarioKey["@key"] = proprietarioMap["@key"]

		id, _ := req["id"].(string)
		quantidade, _ := req["quantidade"].(float64)
		burned, _ := req["burned"].(bool)

		if quantidade <= 0 {
			return nil, errors.WrapError(nil, "A quantidade deve ser maior que zero.")
		}

		tokenMap := make(map[string]interface{})
		tokenMap["@assetType"] = "Token"
		tokenMap["id"] = id
		tokenMap["proprietario"] = updatedProprietarioKey
		tokenMap["quantidade"] = quantidade
		tokenMap["burned"] = burned

		tokenAsset, err := assets.NewAsset(tokenMap)
		if err != nil {
			return nil, errors.WrapError(err, "Falha ao criar ativo 'token'.")
		}

		_, err = tokenAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Erro ao salvar ativo na blockchain.")
		}

		tokenJSON, nerr := json.Marshal(tokenAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "Falha ao converter ativo para JSON.")
		}

		return tokenJSON, nil
	},
}

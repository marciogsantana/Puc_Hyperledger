package assettypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

var Proprietario = assets.AssetType{
	Tag:         "proprietario",
	Label:       "proprietario",
	Description: "proprietario",

	Props: []assets.AssetProp{
		{
			// Primary key
			Required: true,
			IsKey:    true,
			Tag:      "id",
			Label:    "id",
			DataType: "string",                      // Datatypes are identified at datatypes folder
			Writers:  []string{`org1MSP`, "orgMSP"}, // This means only org1 can create the asset (others can edit)
		},
		{
			// Mandatory property
			Required: true,
			Tag:      "nome",
			Label:    "nome",
			DataType: "string",
		},
	},
}

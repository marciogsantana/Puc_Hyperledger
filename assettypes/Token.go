package assettypes

import "github.com/goledgerdev/cc-tools/assets"

var Token = assets.AssetType{
	Tag:         "Token",
	Label:       "Token",
	Description: "Token",

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
			Tag:      "proprietario",
			Label:    "proprietario",
			DataType: "->proprietario",
		},

		{
			// Optional property
			Tag:          "quantidade",
			Label:        "quantidade",
			DefaultValue: 0,
			DataType:     "number",
		},

		{
			// Optional property
			Tag:      "burned",
			Label:    "burned",
			DataType: "boolean",
		},
	},
}

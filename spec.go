package algolia

import sdk "github.com/conduitio/conduit-connector-sdk"

func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "algolia",
		Summary:     "A destination connector for Algolia",
		Description: "TBD",
		Version:     "v0.1.0",
		Author:      "Meroxa, Inc.",
		DestinationParams: map[string]sdk.Parameter{
			DestinationConfigAPIKey: {
				Required:    true,
				Description: "The API key for Algolia.",
			},
			DestinationConfigApplicationID: {
				Required:    true,
				Description: "The Application ID for Algolia.",
			},
			DestinationConfigIndexName: {
				Required:    true,
				Description: "The Algolia index where records get written into.",
			},
		},
		SourceParams: nil,
	}
}

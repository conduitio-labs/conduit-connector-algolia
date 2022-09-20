package algolia

import (
	"context"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

const (
	DestinationConfigAPIKey        = "api_key"
	DestinationConfigApplicationID = "application_id"
	DestinationConfigIndexName     = "index_name"
)

type Destination struct {
	sdk.UnimplementedDestination

	config DestinationConfig
	index  *search.Index
}

type DestinationConfig struct {
	APIKey        string
	ApplicationID string
	IndexName     string
}

func NewDestination() sdk.Destination {
	return sdk.DestinationWithMiddleware(&Destination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *Destination) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
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
	}
}

func (d *Destination) Configure(ctx context.Context, cfg map[string]string) error {
	destCfg := DestinationConfig{
		APIKey:        cfg[DestinationConfigAPIKey],
		ApplicationID: cfg[DestinationConfigApplicationID],
		IndexName:     cfg[DestinationConfigIndexName],
	}

	if destCfg.APIKey == "" {
		return fmt.Errorf("%q is a required parameter", DestinationConfigAPIKey)
	}
	if destCfg.ApplicationID == "" {
		return fmt.Errorf("%q is a required parameter", DestinationConfigApplicationID)
	}
	if destCfg.IndexName == "" {
		return fmt.Errorf("%q is a required parameter", DestinationConfigIndexName)
	}

	d.config = destCfg
	return nil
}

func (d *Destination) Open(ctx context.Context) error {
	client := search.NewClient(d.config.ApplicationID, d.config.APIKey)
	index := client.InitIndex(d.config.IndexName)
	d.index = index
	return nil
}

func (d *Destination) Write(ctx context.Context, records []sdk.Record) (int, error) {
	objects := make([]Object, len(records))
	for i, r := range records {
		objects[i] = Object(r)
	}

	res, err := d.index.SaveObjects(objects, opt.AutoGenerateObjectIDIfNotExist(true))
	if err != nil {
		return 0, fmt.Errorf("could not save objects: %w", err)
	}

	sdk.Logger(ctx).Debug().
		Strs("objectIds", res.ObjectIDs()).
		Msg("saved objects")

	return len(records), nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	// do nothing
	return nil
}

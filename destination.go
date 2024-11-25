package algolia

import (
	"context"
	"fmt"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/opt"
	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/conduitio/conduit-commons/config"
	"github.com/conduitio/conduit-commons/opencdc"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Destination struct {
	sdk.UnimplementedDestination

	config DestinationConfig
	index  *search.Index
}

//go:generate paramgen -output=paramgen.go DestinationConfig
type DestinationConfig struct {
	// API key for Algolia APIs.
	APIKey string `json:"apiKey" validate:"required"`
	// Application ID for Algolia.
	ApplicationID string `json:"applicationID" validate:"required"`
	// IndexName is the name of the index where records get written into.
	IndexName string `json:"indexName" validate:"required"`
}

func NewDestination() sdk.Destination {
	return sdk.DestinationWithMiddleware(&Destination{}, sdk.DefaultDestinationMiddleware()...)
}

func (d *Destination) Configure(ctx context.Context, cfg config.Config) error {
	destCfg := DestinationConfig{
		APIKey:        cfg[DestinationConfigApiKey],
		ApplicationID: cfg[DestinationConfigApplicationID],
		IndexName:     cfg[DestinationConfigIndexName],
	}

	if destCfg.APIKey == "" {
		return fmt.Errorf("%q is a required parameter", DestinationConfigApiKey)
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

func (d *Destination) Write(ctx context.Context, records []opencdc.Record) (int, error) {
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

package algolia

import (
	"context"
	"fmt"

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
	return &Destination{}
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

func (d *Destination) Write(ctx context.Context, record sdk.Record) error {
	res, err := d.index.SaveObject(Object(record))
	if err != nil {
		return fmt.Errorf("could not save object: %w", err)
	}

	sdk.Logger(ctx).Debug().
		Int("taskId", res.TaskID).
		Str("objectId", res.ObjectID).
		Msg("saved object")

	return nil
}

func (d *Destination) Teardown(ctx context.Context) error {
	// do nothing
	return nil
}

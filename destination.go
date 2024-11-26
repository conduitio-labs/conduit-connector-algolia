// Copyright © 2024 Meroxa, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func (d *Destination) Parameters() config.Parameters {
	return d.config.Parameters()
}

func (d *Destination) Configure(_ context.Context, cfg config.Config) error {
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

func (d *Destination) Open(_ context.Context) error {
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

func (d *Destination) Teardown(_ context.Context) error {
	// do nothing
	return nil
}

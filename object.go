// Copyright © 2022 Meroxa, Inc.
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
	"encoding/json"
	"time"

	sdk "github.com/conduitio/conduit-connector-sdk"
)

type Object sdk.Record

func (o Object) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{
		"position":  string(o.Position),
		"key":       parseData(o.Key),
		"payload":   parseData(o.Payload),
		"createdAt": o.CreatedAt.Format(time.RFC3339),
	}
	return json.Marshal(out)
}

func parseData(d sdk.Data) interface{} {
	var out map[string]interface{}
	err := json.Unmarshal(d.Bytes(), &out)
	if err != nil {
		return string(d.Bytes())
	}
	return out
}

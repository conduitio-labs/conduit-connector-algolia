package algolia

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

// version is set during the build process (i.e. the Makefile).
// It follows Go's convention for module version, where the version
// starts with the letter v, followed by a semantic version.
var version = "dev"

func Specification() sdk.Specification {
	return sdk.Specification{
		Name:        "algolia",
		Summary:     "A destination connector for Algolia",
		Description: "TBD",
		Version:     version,
		Author:      "Meroxa, Inc.",
	}
}

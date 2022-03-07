package main

import (
	algolia "github.com/conduitio/conduit-connector-algolia"
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func main() {
	sdk.Serve(
		algolia.Specification,
		nil,
		algolia.NewDestination,
		)
}
# Conduit Connector Algolia

### General

This connector allows you to move data from any [Conduit Source](https://www.conduit.io/docs/connectors/overview) to an [Algolia Index](https://www.algolia.com/doc/guides/sending-and-managing-data/send-and-update-your-data/). This connector is a destination connector.

### How to build it

Run:

```bash
go build -o algolia cmd/algolia/main.go`
```

### How it works

Under the hood, the connector uses [Algolia's Go Library](https://www.algolia.com/doc/api-client/getting-started/install/go/?client=go) to send data to Algolia.

To learn more see: [How to build a Conduit Connector](https://www.conduit.io/guides/build-a-conduit-connector)

### Configuration

There's no global, connector configuration. Each connector instance is configured separately.

| name                 | part of             | description                                                                                                                                                                        | required | default value |
|----------------------|---------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------|---------------|
| `APIKey`            | destination | The [API key](https://www.algolia.com/doc/guides/security/api-keys/) for Algolia.                                                                                                                      | true     |               |
| `ApplicationID`              | destination | The [Application ID](https://www.algolia.com/doc/guides/security/api-keys/) for Algolia.                                                                                                                           | true     |               |
| `IndexName`         | destination | The Algolia index where records get written into.                                                                          |  true   |               |
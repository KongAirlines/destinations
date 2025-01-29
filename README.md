## Destinations API

Provides the KongAir routing information including flight destination airport codes, cities and images.

### Security

See [Security](SECURITY.md) for information on how to report security vulnerabilities.

### Specification

The API specification can be found in the [openapi.yaml](openapi.yaml) file.

### Usage

The repository provides a `Makefile` with common usage.

#### To run unit tests

```
make test
```

#### To build the server

```
make build
```

#### To run the server on the default port

```
make run
```

In the `Makefile`, the default port is read from the `KONG_AIR_DESTINATIONS_PORT`
env var which is loaded via the parent [base.mk](../../base.mk) file.

Alternatively the desired port can be passed to the built server executable directly,
for example:

```sh
./destinations <port>
```

### Example client requests

Get all destinations:
```
curl -s localhost:8081/destinations
```

Get destination by airport code
```
curl -s localhost:8081/destinations/SIN
```


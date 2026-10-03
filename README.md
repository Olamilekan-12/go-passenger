# GoPick: Passenger Service

The passenger-facing backend service for GoPick, a ride-hailing platform.
Handles passenger identity, profiles and ride requests.

Part of the GoPick system:
- **go-passenger** (this repo) : passenger API
- **go-driver** : driver API
- **go-infra** : shared infrastructure and architecture docs

## Tech stack

- **Go** : service implementation
- **PostgreSQL** : primary datastore

## Status

Early development. Nothing is stable yet.

## Getting started

Requirements: Go 1.26+, PostgreSQL 16+

```bash
git clone git@github.com:Olamilekan-12/go-passenger.git
cd go-passenger
cp .env.example .env   # then fill in your own values
```

## License

MIT


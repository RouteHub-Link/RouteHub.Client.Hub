# Routehub.Client.HUB

[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/routehub-helm)](https://artifacthub.io/packages/search?repo=routehub-helm)

This project is HUB application aka client. Planned to deployed for every domain for user project.

Has two hosting modes;

- MQTT
- REST

MQTT Used for receiving updates from GraphQL Service.

Please check out [RouteHub-Link/RouteHub.Service.GraphQL](https://github.com/RouteHub-Link/RouteHub.Service.GraphQL "RouteHub GraphQL Service").

## Development

Template engine is templ. you need to install it before running the project and you need to run the following command for changes to take effect.

```bash
- go install github.com/a-h/templ/cmd/templ@latest
- go get github.com/a-h/templ
```

Please check out the Makefile for commands.

```bash
make serve-rest
make serve-mqtt
make hotserve
```

Please use .env file for required environment variables. For easy development SEED is provided. Check .env.development file. Copy it to .env file.
For development you just need to connect 2 services. Redis and TimeScaleDB. You can use docker-compose file for that or you can run them by yourself. Pleaase check Makefile for commands.

```bash
make keydb
make timescaledb
```

If you use podman you can use podman- as prefix. For example;

```bash
make podman-keydb
```

## Deployment Requirements

- REST as self (for handling request's)
- MQTT as self (for handling update's & sending reports)
- Redis as keydb (keydb not required but suggested. Used as database)
- TimeScaleDB (Storing request's for analytics)

## Deployment

Please check out the helm chart for deployment. [RouteHub-Link/RouteHub.Helm](https://github.com/RouteHub-Link/RouteHub.HELM/blob/main/charts/routehub-server/readme.md)

## Technologies

- GO
- TimeScaleDB
- Redis as KeyDB
- MQTT
- Rest Echo
- Template Rendering templ
- Bulma CSS
- Docker

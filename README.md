# kafkactl

## Table of contents

- [kafkactl](#kafkactl)
  - [Table of contents](#table-of-contents)
  - [Overview](#overview)
  - [Build](#build)
  - [Development](#development)

## Overview

`kafkactl` is a CLI tool to interact with Kafka through the Confluent Kafka Rest Proxy.

The api package follows the [OpenAPI spec](https://github.com/confluentinc/kafka-rest/blob/v6.2.1/api/v3/openapi.yaml) of Kafka Rest Proxy.

## Docs

Documentation for the CLI (auto-generated by Cobra) is available [here](./docs/kafkactl.md).

### Generate docs

To generate CLI markup docs (provided by Cobra), run the following command at the root of the repo

```bash
make gendoc
```

## Roadmap

- [ ] Add "Consume records" command
- [ ] Add "Produce records" command
- [ ] Add get broker tasks command
- [x] Add create command
  - [x] Add create acl
- [x] Add update command
  - [x] Add update broker-config
  - [x] Add update cluster-config
  - [x] Add update topic-config
- [x] Add delete command
  - [x] Add delete acl
  - [ ] Add delete topic
- [ ] Add describe command
  - [ ] Add describe cluster
  - [ ] Add describe topic
  - [ ] Add describe consumer-group

## Build

> Output binary will be copied to repo/bin directory

Build and copy `kafkactl` into the `bin` directory

```bash
make genbin
```

Or build for another platform

```bash
# macOS with Apple Silicon
make genbin GOOS=darwin GOARCH=arm64
```

## Development

> This requires the docker engine to be running and docker-compose installed

Start a local Kafka cluster + Kafka Rest Proxy

```bash
make dev-cluster-start
```

Tail the logs of the local Kafka Rest Proxy

```bash
make dev-cluster-logs
```

Stop a local Kafka cluster + Kafka Rest Proxy

```bash
make dev-cluster-stop
```

Destroy the local Kafka cluster + Kafka Rest Proxy

```bash
make dev-cluster-down
```

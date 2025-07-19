# go-nitrado

[![Build Test Release](https://github.com/danstis/go-nitrado/actions/workflows/build.yml/badge.svg)](https://github.com/danstis/go-nitrado/actions/workflows/build.yml)
[![DeepSource](https://deepsource.io/gh/danstis/go-nitrado.svg/?label=active+issues)](https://deepsource.io/gh/danstis/go-nitrado/?ref=repository-badge)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=danstis_go-nitrado&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=danstis_go-nitrado)
[![codecov](https://codecov.io/gh/danstis/go-nitrado/branch/master/graph/badge.svg?token=Q2T27EQ2XM)](https://codecov.io/gh/danstis/go-nitrado)

Go library for accessing the [nitrado.net](https://doc.nitrado.net/) API.

**Note:** go-nitrado is currently in development, so its API may have breaking changes.

## Usage

```go
import "github.com/danstis/go-nitrado/nitrado"
```

Create a new Nitrado client instance, then use provided methods on the client to
access the API. For example, to list all services:

```go
client := nitrado.NewClient("YourNitradoToken")
services, resp, err := client.Services.List()
```

## Feature requests

Feature request tracking and voting is being tracked using [GitHub discussions](https://github.com/danstis/go-openxbl/discussions/categories/ideas).

## Credit

This API Client is based on the format and hard work of the [go-github client library](https://github.com/google/go-github).

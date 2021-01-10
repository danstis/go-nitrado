# go-nitrado

[![Build Test and Release](https://github.com/danstis/go-nitrado/workflows/Build%20Test%20and%20Release/badge.svg)](https://github.com/danstis/go-nitrado/actions?query=workflow%3A%22Build+Test+and+Release%22)
[![DeepSource](https://deepsource.io/gh/danstis/go-nitrado.svg/?label=active+issues)](https://deepsource.io/gh/danstis/go-nitrado/?ref=repository-badge)
[![Maintainability](https://api.codeclimate.com/v1/badges/8a7e6990547ee87d167a/maintainability)](https://codeclimate.com/github/danstis/go-nitrado/maintainability)
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

Feature request tracking and voting is being tracked by feathub:
[![Feature Requests](https://feathub.com/danstis/go-nitrado?format=svg)](https://feathub.com/danstis/go-nitrado)

## Credit

This API Client is based on the format and hard work of the [go-github client library](https://github.com/google/go-github).

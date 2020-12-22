# go-nitrado

[![Build Test and Release](https://github.com/danstis/go-nitrado/workflows/Build%20Test%20and%20Release/badge.svg)](https://github.com/danstis/go-nitrado/actions?query=workflow%3A%22Build+Test+and+Release%22)
[![DeepSource](https://deepsource.io/gh/danstis/go-nitrado.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/danstis/go-nitrado/?ref=repository-badge)
[![codecov](https://codecov.io/gh/danstis/go-nitrado/branch/master/graph/badge.svg?token=Q2T27EQ2XM)](https://codecov.io/gh/danstis/go-nitrado)

Go library for accessing the [nitrado.net](https://doc.nitrado.net/) API.

**Note:** go-nitrado is currently in development, so its API may have slightly breaking changes if we find better ways of doing things.

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

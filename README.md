# go-nitrado

![Build Test and Release](https://github.com/danstis/go-nitrado/workflows/Build%20Test%20and%20Release/badge.svg)
[![DeepSource](https://deepsource.io/gh/danstis/go-nitrado.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/danstis/go-nitrado/?ref=repository-badge)

Go library for accessing the Nitrado.net API.

**Note:** go-nitrado is currently in development, so its API may have slightly breaking changes if we find better ways of doing things.

### Usage ###

```go
import "github.com/danstis/go-nitrado/nitrado"
```

Create a new Nitrado client instance, then use provided methods on the client to
access the API. For example, to list all workspaces:

```go
client := nitrado.NewClient(nil)
services, err := client.ListServices()
```

### Authentication ###

The go-nitrado library does not directly handle authentication. Instead, when
creating a new client, pass an `http.Client` that can handle authentication for
you. The easiest way to do this is using the [goauth2][] library, but you can
always use any other library that provides an `http.Client`. If you have an OAuth2
access token, you can use it with the goauth2 using:

```go
t := &oauth.Transport{
  Token: &oauth.Token{AccessToken: "... your access token ..."},
}

client := nitrado.NewClient(t.Client())

// List all projects for the authenticated user
projects, err := client.ListProjects(opt)
```

See the [goauth2 docs][] for complete instructions on using that library.

[goauth2]: https://github.com/golang/oauth2
[goauth2 docs]: https://godoc.org/golang.org/x/oauth2

package nitrado

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURI string        = "https://api.nitrado.net/"
	retryCount     int           = 10
	retryDelay     time.Duration = 2 * time.Second
	userAgent      string        = "go-nitrado"
)

// Client represents the config of the Nitrado.net Client
type Client struct {
	sync.Mutex
	client    *http.Client
	token     string
	UserAgent string

	// Base URL for API requests. Defaults to the public GitHub API, but can be
	// set to a domain endpoint to use with GitHub Enterprise. BaseURL should
	// always be specified with a trailing slash.
	BaseURI *url.URL

	common apiService // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Nitrado API.
	FileServerService   *FileServerService
	GameServers         *GameServersService
	GameServersSettings *GSSettingsService
	Services            *ServicesService
	GameServerStats     *GameServerStatsService
}

type apiService struct {
	client *Client
}

// addOptions adds the parameters in opts as URL query parameters to s. opts
// must be a struct whose fields may contain "url" tags.
func addOptions(s string, opts interface{}) (string, error) {
	v := reflect.ValueOf(opts)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return s, nil
	}

	u, err := url.Parse(s)
	if err != nil {
		return s, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return s, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURI of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURI.Path, "/") {
		return nil, fmt.Errorf("BaseURI must have a trailing slash, but %q does not", c.BaseURI)
	}
	u, err := c.BaseURI.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do sends an API request to the Nitrado API and returns a response object.
// The API response is JSON decoded and stored in the value pointed to by v,
// or returned as an error if an API error has occurred. If v implements the
// io.Writer interface, the raw response body will be written to v, without
// attempting to first decode it.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {

	// Do the request
	var err error = nil
	var resp *http.Response
	for i := 0; i < retryCount; i++ {
		resp, err = c.client.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			time.Sleep(retryDelay)
		}
	}
	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// NewClient creates a new instance of a NitradoAPI
func NewClient(apiToken string) *Client {
	baseURL, _ := url.Parse(defaultBaseURI)

	c := &Client{
		BaseURI:   baseURL,
		token:     apiToken,
		client:    &http.Client{},
		UserAgent: userAgent,
	}

	c.common.client = c

	c.FileServerService = (*FileServerService)(&c.common)
	c.GameServers = (*GameServersService)(&c.common)
	c.GameServersSettings = (*GSSettingsService)(&c.common)
	c.Services = (*ServicesService)(&c.common)
	c.GameServerStats = (*GameServerStatsService)(&c.common)

	return c
}

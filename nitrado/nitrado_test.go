package nitrado

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

const (
	// baseURLPath is a non-empty Client.BaseURL path to use during tests,
	// to ensure relative URLs are used for all endpoints.
	baseURLPath = "/api"
	token       = "abcdefg_1234567.abcdef"
)

// setup sets up a test HTTP server along with a Nitrado.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// We want to ensure that tests catch mistakes where the endpoint URL is
	// specified as absolute rather than relative. It only makes a difference
	// when there's a non-empty base URL path. So, use that.
	apiHandler := http.NewServeMux()
	apiHandler.Handle(baseURLPath+"/", http.StripPrefix(baseURLPath, mux))
	apiHandler.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(os.Stderr, "FAIL: Client.BaseURL path prefix is not preserved in the request URL:")
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\t"+req.URL.String())
		fmt.Fprintln(os.Stderr)
		fmt.Fprintln(os.Stderr, "\tDid you accidentally use an absolute endpoint URL rather than relative?")
		http.Error(w, "Client.BaseURL path prefix is not preserved in the request URL.", http.StatusInternalServerError)
	})

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the Nitrado client being tested and is
	// configured to use test server.
	client = NewClient(token)
	url, _ := url.Parse(server.URL + baseURLPath + "/")
	client.BaseURI = url

	return client, mux, server.URL, server.Close
}

// testMethod checks the request's method.
func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

// testURLParseError checks for errors in Parsing the URL.
func testURLParseError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

// TestNewClient tests the NewClient() method.
func TestNewClient(t *testing.T) {
	c := NewClient(token)

	if got, want := c.BaseURI.String(), defaultBaseURI; got != want {
		t.Errorf("NewClient BaseURI is %v, want %v", got, want)
	}
	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}

	c2 := NewClient(token)
	if c.client == c2.client {
		t.Error("NewClient returned same http.Clients, but they should differ")
	}
}

// TestNewRequest tests the NewRequest() method using various input.
func TestNewRequest(t *testing.T) {
	c := NewClient(token)

	inURL, outURL := "/foo", defaultBaseURI+"foo"
	inBody, outBody := &[]string{"Test", "Test2"}, `["Test","Test2"]`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// Test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// Test that body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}

	// Test that default user-agent is attached to the request
	if got, want := req.Header.Get("User-Agent"), c.UserAgent; got != want {
		t.Errorf("NewRequest() User-Agent is %v, want %v", got, want)
	}

	// Test a BaseURI without a trailing slash
	c.BaseURI, _ = url.Parse("http://url.invalid")
	wantErr := fmt.Errorf("BaseURI must have a trailing slash, but %q does not", c.BaseURI)
	_, gotErr := c.NewRequest("GET", inURL, inBody)
	if gotErr.Error() != wantErr.Error() {
		t.Errorf("NewRequest error is %#v, wanted %#v", gotErr, wantErr)
	}
}

// TestNewRequest_invalidJSON tests the NewRequest() method with invalid JSON as the input object.
func TestNewRequest_invalidJSON(t *testing.T) {
	c := NewClient(token)

	type T struct {
		A map[interface{}]interface{}
	}
	_, err := c.NewRequest("GET", ".", &T{})

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*json.UnsupportedTypeError); !ok {
		t.Errorf("Expected a JSON error; got %#v.", err)
	}
}

// TestNewRequest_badURL tests the NewRequest() method with an invalid URL.
func TestNewRequest_badURL(t *testing.T) {
	c := NewClient(token)
	_, err := c.NewRequest("GET", ":", nil)
	testURLParseError(t, err)
}

// TestNewRequest_badMethod tests the NewRequest() method with a request with an invalid method.
func TestNewRequest_badMethod(t *testing.T) {
	c := NewClient(token)
	if _, err := c.NewRequest("BOGUS\nMETHOD", ".", nil); err == nil {
		t.Fatal("NewRequest returned nil; expected error")
	}
}

// Test_addOptions_empty tests the addOptions() function with an empty opts parameter.
func Test_addOptions_empty(t *testing.T) {
	if _, err := addOptions("test/path", ""); err == nil {
		t.Fatal("NewRequest returned nil; expected error")
	}
}

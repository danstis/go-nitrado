package nitrado

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
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
	Services *ServicesService
}

type apiService struct {
	client *Client
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
	c.Lock()
	defer c.Unlock()

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

	c.Services = (*ServicesService)(&c.common)

	return c
}

// // Services gets a list of services on the account
// func (s *Client) Services() ([]Service, error) {
// 	var result []Service
// 	// No filter given, return all services
// 	result, err := serviceList(s)
// 	if err != nil {
// 		return []Service{}, err
// 	}
// 	return result, nil
// }

// // Service returns an individual service by ID
// func (s *Client) Service(id int) (Service, error) {
// 	var result Service
// 	// Return a filtered list of services based on the ID given
// 	svc, err := serviceList(s)
// 	if err != nil {
// 		return Service{}, err
// 	}
// 	for i := range svc {
// 		if svc[i].ID == id {
// 			result = svc[i]
// 		}
// 	}

// 	return result, nil
// }

// // LogFiles gets a list of DayZ log files on the gameserver
// func (s *Client) LogFiles(svc Service) ([]File, error) {
// 	var result []File
// 	// No filter given, return all logs
// 	result, err := logFiles(s, svc)
// 	if err != nil {
// 		return []File{}, err
// 	}

// 	sort.Slice(result, func(i, j int) bool {
// 		return result[i].ModifiedAt < result[j].ModifiedAt
// 	})

// 	return result, nil
// }

// // GameServer returns an individual service by ID
// func (s *Client) GameServer(svcID int) (GameServer, error) {
// 	gs, err := gameServerDetails(s, svcID)
// 	if err != nil {
// 		return GameServer{}, err
// 	}

// 	return gs, nil
// }

// // DownloadLink provides a download link for a specified file
// func (s *Client) DownloadLink(svc Service, file File) (string, error) {
// 	fdl, err := link(s, svc, file)
// 	if err != nil {
// 		return "", err
// 	}

// 	return fdl, nil
// }

// // Restart a service
// func (s *Client) Restart(svcID int) error {
// 	resp, err := s.apiRequest("POST", fmt.Sprintf("%s/services/%v/gameservers/restart", s.BaseURI, svcID), s.token, nil)
// 	if err != nil {
// 		return fmt.Errorf("Nitrado API query failed: %v", err)
// 	}
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("failed to restart game server: %v", resp.Status)
// 	}
// 	return nil
// }

// func (s *Client) apiRequest(method, url, token string, body io.Reader) (*http.Response, error) {
// 	s.Lock()
// 	defer s.Unlock()
// 	req, err := http.NewRequest(method, url, body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create new request, %v", err)
// 	}
// 	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

// 	// Do the request
// 	for i := 0; i < retryCount; i++ {
// 		resp, err := s.client.Do(req)
// 		if err != nil || resp.StatusCode >= 400 {
// 			log.Printf("Request failed retying in %s...\nError: %v\nResponse: %v", retryDelay, err, resp)
// 			time.Sleep(retryDelay)
// 		} else {
// 			return resp, nil
// 		}
// 	}

// 	return nil, fmt.Errorf("request retries exceeded try again shortly")
// }

// func serviceList(s *Client) ([]Service, error) {
// 	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services", s.BaseURI), s.token, nil)
// 	if err != nil {
// 		return []Service{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return []Service{}, fmt.Errorf(resp.Status)
// 	}

// 	var serviceList ServiceList

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return []Service{}, fmt.Errorf("error reading response: %v", err)
// 	}
// 	err = json.Unmarshal(body, &serviceList)
// 	if err != nil {
// 		return []Service{}, fmt.Errorf("error parsing JSON response: %v", err)
// 	}

// 	return serviceList.Data.Services, nil
// }

// func gameServerDetails(s *Client, svcID int) (GameServer, error) {
// 	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers", s.BaseURI, svcID), s.token, nil)
// 	if err != nil {
// 		return GameServer{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return GameServer{}, fmt.Errorf(resp.Status)
// 	}

// 	var gsResp GameServerQuery

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return GameServer{}, fmt.Errorf("error reading response: %v", err)
// 	}
// 	err = json.Unmarshal(body, &gsResp)
// 	if err != nil {
// 		return GameServer{}, fmt.Errorf("error parsing JSON response: %v", err)
// 	}

// 	return gsResp.Data.GameServer, nil
// }

// func logFiles(s *Client, svc Service) ([]File, error) {
// 	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers/file_server/list?dir=/games/%s/noftp/dayzxb/config", s.BaseURI, svc.ID, svc.Username), s.token, nil)
// 	if err != nil {
// 		return []File{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return []File{}, fmt.Errorf(resp.Status)
// 	}

// 	var fileList FileList

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return []File{}, fmt.Errorf("error reading response: %v", err)
// 	}
// 	err = json.Unmarshal(body, &fileList)
// 	if err != nil {
// 		return []File{}, fmt.Errorf("error parsing JSON response: %v", err)
// 	}

// 	return fileList.Data.Entries, nil
// }

// func link(s *Client, svc Service, file File) (string, error) {
// 	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers/file_server/download?file=%s", s.BaseURI, svc.ID, file.Path), s.token, nil)
// 	if err != nil {
// 		return "", fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return "", fmt.Errorf(resp.Status)
// 	}

// 	var link FileLink

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading response: %v", err)
// 	}
// 	err = json.Unmarshal(body, &link)
// 	if err != nil {
// 		return "", fmt.Errorf("error parsing JSON response: %v", err)
// 	}

// 	return link.Data.Token.URL, nil
// }

package nitrado

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	// APIBASEURL contains the base URL of the Nitrado.net API
	APIBASEURL string = "https://api.nitrado.net"
	// RETRIES controls the number or retries that will be attempted if the HTTP request fails
	RETRIES int = 10
	// RETRYDELAY controls the delay in seconds between retries
	RETRYDELAY time.Duration = 15 * time.Second
)

// API represents the config of the Nitrado.net API
type API struct {
	sync.Mutex
	BaseURI string
	Client  *http.Client
	Token   string
}

// New creates a new instance of a NitradoAPI
func New(t string) *API {
	server := API{}
	server.BaseURI = APIBASEURL
	server.Token = t
	server.Client = &http.Client{}

	return &server
}

// Services gets a list of services on the account
func (s *API) Services() ([]Service, error) {
	var result []Service
	// No filter given, return all services
	result, err := serviceList(s)
	if err != nil {
		return []Service{}, err
	}
	return result, nil
}

// Service returns an individual service by ID
func (s *API) Service(id int) (Service, error) {
	var result Service
	// Return a filtered list of services based on the ID given
	svc, err := serviceList(s)
	if err != nil {
		return Service{}, err
	}
	for i := range svc {
		if svc[i].ID == id {
			result = svc[i]
		}
	}

	return result, nil
}

// LogFiles gets a list of DayZ log files on the gameserver
func (s *API) LogFiles(svc Service) ([]File, error) {
	var result []File
	// No filter given, return all logs
	result, err := logFiles(s, svc)
	if err != nil {
		return []File{}, err
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].ModifiedAt < result[j].ModifiedAt
	})

	return result, nil
}

// GameServer returns an individual service by ID
func (s *API) GameServer(svcID int) (GameServer, error) {
	gs, err := gameServerDetails(s, svcID)
	if err != nil {
		return GameServer{}, err
	}

	return gs, nil
}

// DownloadLink provides a download link for a specified file
func (s *API) DownloadLink(svc Service, file File) (string, error) {
	fdl, err := link(s, svc, file)
	if err != nil {
		return "", err
	}

	return fdl, nil
}

// Restart a service
func (s *API) Restart(svcID int) error {
	resp, err := s.apiRequest("POST", fmt.Sprintf("%s/services/%v/gameservers/restart", s.BaseURI, svcID), s.Token, nil)
	if err != nil {
		return fmt.Errorf("Nitrado API query failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to restart game server: %v", resp.Status)
	}
	return nil
}

func (s *API) apiRequest(method, url, token string, body io.Reader) (*http.Response, error) {
	s.Lock()
	defer s.Unlock()
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request, %v", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	// Do the request
	for i := 0; i < RETRIES; i++ {
		resp, err := s.Client.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			log.Printf("Request failed retying in %s...\nError: %v\nResponse: %v", RETRYDELAY, err, resp)
			time.Sleep(RETRYDELAY)
		} else {
			return resp, nil
		}
	}

	return nil, fmt.Errorf("request retries exceeded try again shortly")
}

func serviceList(s *API) ([]Service, error) {
	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services", s.BaseURI), s.Token, nil)
	if err != nil {
		return []Service{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Service{}, fmt.Errorf(resp.Status)
	}

	var serviceList Services

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Service{}, fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, &serviceList)
	if err != nil {
		return []Service{}, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return serviceList.Data.Services, nil
}

func gameServerDetails(s *API, svcID int) (GameServer, error) {
	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers", s.BaseURI, svcID), s.Token, nil)
	if err != nil {
		return GameServer{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return GameServer{}, fmt.Errorf(resp.Status)
	}

	var gsResp GameServerQuery

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return GameServer{}, fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, &gsResp)
	if err != nil {
		return GameServer{}, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return gsResp.Data.GameServer, nil
}

func logFiles(s *API, svc Service) ([]File, error) {
	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers/file_server/list?dir=/games/%s/noftp/dayzxb/config", s.BaseURI, svc.ID, svc.Username), s.Token, nil)
	if err != nil {
		return []File{}, fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []File{}, fmt.Errorf(resp.Status)
	}

	var fileList FileList

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []File{}, fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, &fileList)
	if err != nil {
		return []File{}, fmt.Errorf("error parsing JSON response: %v", err)
	}

	return fileList.Data.Entries, nil
}

func link(s *API, svc Service, file File) (string, error) {
	resp, err := s.apiRequest("GET", fmt.Sprintf("%s/services/%v/gameservers/file_server/download?file=%s", s.BaseURI, svc.ID, file.Path), s.Token, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}

	var link FileLink

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}
	err = json.Unmarshal(body, &link)
	if err != nil {
		return "", fmt.Errorf("error parsing JSON response: %v", err)
	}

	return link.Data.Token.URL, nil
}

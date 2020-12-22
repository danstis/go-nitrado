package nitrado

import (
	"fmt"
	"net/http"
)

// Generated structs from https://mholt.github.io/json-to-go/

// ServicesService provides access to the service related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/
type ServicesService apiService

// Service contains the structure of a service object
type Service struct {
	ID                    int         `json:"id"`
	LocationID            int         `json:"location_id"`
	Status                string      `json:"status"`
	WebsocketToken        string      `json:"websocket_token"`
	UserID                int         `json:"user_id"`
	Comment               interface{} `json:"comment"`
	AutoExtension         bool        `json:"auto_extension"`
	AutoExtensionDuration int         `json:"auto_extension_duration"`
	Type                  string      `json:"type"`
	TypeHuman             string      `json:"type_human"`
	Details               struct {
		Address       string `json:"address"`
		Name          string `json:"name"`
		Game          string `json:"game"`
		PortlistShort string `json:"portlist_short"`
		FolderShort   string `json:"folder_short"`
		Slots         int    `json:"slots"`
	} `json:"details"`
	StartDate    string   `json:"start_date"`
	SuspendDate  string   `json:"suspend_date"`
	DeleteDate   string   `json:"delete_date"`
	SuspendingIn int      `json:"suspending_in"`
	DeletingIn   int      `json:"deleting_in"`
	Username     string   `json:"username"`
	Roles        []string `json:"roles"`
}

// ServiceListResp contains a list of services
type ServiceListResp struct {
	Status string `json:"status"`
	Data   struct {
		Services []Service `json:"services"`
	} `json:"data"`
}

// ServiceDetailResp contains a list of services
type ServiceDetailResp struct {
	Status string `json:"status"`
	Data   struct {
		Service Service `json:"service"`
	} `json:"data"`
}

// List all services.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Service-List
func (s *ServicesService) List() (*[]Service, *http.Response, error) {
	var services *[]Service
	req, err := s.client.NewRequest("GET", "services", nil)
	if err != nil {
		return services, nil, err
	}

	var serviceListResp *ServiceListResp
	resp, err := s.client.Do(req, &serviceListResp)
	if err != nil {
		return services, resp, err
	}

	services = &serviceListResp.Data.Services

	return services, resp, nil
}

// Get a Service by ID.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Service-Details
func (s *ServicesService) Get(ID int) (*Service, *http.Response, error) {
	u := fmt.Sprintf("services/%v", ID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var serviceDetailResp *ServiceDetailResp
	resp, err := s.client.Do(req, &serviceDetailResp)
	if err != nil {
		return nil, resp, err
	}

	return &serviceDetailResp.Data.Service, resp, nil
}

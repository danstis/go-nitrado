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
	ID                    int         `json:"id,omitempty"`
	LocationID            int         `json:"location_id,omitempty"`
	Status                string      `json:"status,omitempty"`
	WebsocketToken        string      `json:"websocket_token,omitempty"`
	UserID                int         `json:"user_id,omitempty"`
	Comment               interface{} `json:"comment,omitempty"`
	AutoExtension         bool        `json:"auto_extension,omitempty"`
	AutoExtensionDuration int         `json:"auto_extension_duration,omitempty"`
	Type                  string      `json:"type,omitempty"`
	TypeHuman             string      `json:"type_human,omitempty"`
	Details               struct {
		Address       string `json:"address,omitempty"`
		Name          string `json:"name,omitempty"`
		Game          string `json:"game,omitempty"`
		PortlistShort string `json:"portlist_short,omitempty"`
		FolderShort   string `json:"folder_short,omitempty"`
		Slots         int    `json:"slots,omitempty"`
	} `json:"details,omitempty"`
	StartDate    string   `json:"start_date,omitempty"`
	SuspendDate  string   `json:"suspend_date,omitempty"`
	DeleteDate   string   `json:"delete_date,omitempty"`
	SuspendingIn int      `json:"suspending_in,omitempty"`
	DeletingIn   int      `json:"deleting_in,omitempty"`
	Username     string   `json:"username,omitempty"`
	Roles        []string `json:"roles,omitempty"`
}

// ServiceListResp contains a list of services
type ServiceListResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Services []Service `json:"services,omitempty"`
	} `json:"data,omitempty"`
}

// ServiceDetailResp contains a list of services
type ServiceDetailResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Service Service `json:"service,omitempty"`
	} `json:"data,omitempty"`
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

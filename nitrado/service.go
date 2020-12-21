package nitrado

import (
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

// ServiceList contains a list of services
type ServiceList struct {
	Status string `json:"status"`
	Data   struct {
		Services []Service `json:"services"`
	} `json:"data"`
}

// List all services.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Service-List
func (s *ServicesService) List() (*ServiceList, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "services", nil)
	if err != nil {
		return &ServiceList{}, nil, err
	}

	var services *ServiceList
	resp, err := s.client.Do(req, &services)
	if err != nil {
		return services, resp, err
	}

	return services, resp, nil
}

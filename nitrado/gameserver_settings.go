package nitrado

import (
	"fmt"
)

// Generated structs from https://mholt.github.io/json-to-go/

// GSSettingsService provides access to the service related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/
type GSSettingsService apiService

// GSSettingsResp contains the response object from the download method on a fileserver
type GSSettingsResp struct {
	Status string `json:"status,omitempty"`
}

// GSSettingsUpdateOptions controls the query string settings that a settings request can take.
type GSSettingsUpdateOptions struct {
	Category string `url:"category,omitempty"`
	Key      string `url:"key,omitempty"`
	Value    string `url:"value"`
}

// Update a setting on a GameServer by service ID.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-Details
// Requires a settings category as well as a key and value for the setting.
func (s *GSSettingsService) Update(serviceID int, opts GSSettingsUpdateOptions) error {
	if opts.Category == "" || opts.Key == "" {
		return fmt.Errorf("category and key must not be blank. category=%q, key=%q", opts.Category, opts.Key)
	}
	u := fmt.Sprintf("services/%v/gameservers/settings", serviceID)
	u, err := addOptions(u, opts)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return err
	}

	var settingsResp *GSSettingsResp
	_, err = s.client.Do(req, &settingsResp)
	if err != nil {
		return err
	}
	if settingsResp.Status != "success" {
		return fmt.Errorf("status %q", settingsResp.Status)
	}

	return nil
}

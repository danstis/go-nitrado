package nitrado

import (
	"fmt"
	"net/http"
)

// Generated structs from https://mholt.github.io/json-to-go/

// GameServersService provides access to the service related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/
type GameServersService apiService

// GameServer contains a Game Server details
type GameServer struct {
	Status           string `json:"status,omitempty"`
	LastStatusChange int    `json:"last_status_change,omitempty"`
	MustBeStarted    bool   `json:"must_be_started,omitempty"`
	Username         string `json:"username,omitempty"`
	UserID           int    `json:"user_id,omitempty"`
	ServiceID        int    `json:"service_id,omitempty"`
	IP               string `json:"ip,omitempty"`
	Port             int    `json:"port,omitempty"`
	QueryPort        int    `json:"query_port,omitempty"`
	RconPort         int    `json:"rcon_port,omitempty"`
	Type             string `json:"type,omitempty"`
	Memory           string `json:"memory,omitempty"`
	MemoryMb         int    `json:"memory_mb,omitempty"`
	Game             string `json:"game,omitempty"`
	GameHuman        string `json:"game_human,omitempty"`
	GameSpecific     struct {
		Path         string   `json:"path,omitempty"`
		UpdateStatus string   `json:"update_status,omitempty"`
		LastUpdate   string   `json:"last_update,omitempty"`
		LogFiles     []string `json:"log_files,omitempty"`
		ConfigFiles  []string `json:"config_files,omitempty"`
	} `json:"game_specific,omitempty"`
	Slots       int    `json:"slots,omitempty"`
	Location    string `json:"location,omitempty"`
	Credentials struct {
		Ftp struct {
			Hostname string `json:"hostname,omitempty"`
			Port     int    `json:"port,omitempty"`
			Username string `json:"username,omitempty"`
			Password string `json:"password,omitempty"`
		} `json:"ftp,omitempty"`
		Mysql struct {
			Hostname string `json:"hostname,omitempty"`
			Port     int    `json:"port,omitempty"`
			Username string `json:"username,omitempty"`
			Password string `json:"password,omitempty"`
			Database string `json:"database,omitempty"`
		} `json:"mysql,omitempty"`
	} `json:"credentials,omitempty"`
	Settings struct {
		Config struct {
			Hostname                    string `json:"hostname,omitempty"`
			VonCodecQuality             string `json:"vonCodecQuality,omitempty"`
			DisableVoN                  string `json:"disableVoN,omitempty"`
			Password                    string `json:"password,omitempty"`
			ServerTimeAcceleration      string `json:"serverTimeAcceleration,omitempty"`
			ServerNightTimeAcceleration string `json:"serverNightTimeAcceleration,omitempty"`
			ServerTimePersistent        string `json:"serverTimePersistent,omitempty"`
			Disable3RdPerson            string `json:"disable3rdPerson,omitempty"`
			DisableCrosshair            string `json:"disableCrosshair,omitempty"`
			UseServerTime               string `json:"useServerTime,omitempty"`
			CustomServerTime            string `json:"customServerTime,omitempty"`
			EnableMouseAndKeyboard      string `json:"enableMouseAndKeyboard,omitempty"`
			EnableWhitelist             string `json:"enableWhitelist,omitempty"`
			Mission                     string `json:"mission,omitempty"`
			AdminLogPlayerHitsOnly      string `json:"adminLogPlayerHitsOnly,omitempty"`
			AdminLogPlacement           string `json:"adminLogPlacement,omitempty"`
			AdminLogBuildActions        string `json:"adminLogBuildActions,omitempty"`
			AdminLogPlayerList          string `json:"adminLogPlayerList,omitempty"`
			LightingConfig              string `json:"lightingConfig,omitempty"`
			DisablePersonalLight        string `json:"disablePersonalLight,omitempty"`
			DisableBaseDamage           string `json:"disableBaseDamage,omitempty"`
			DisableContainerDamage      string `json:"disableContainerDamage,omitempty"`
		} `json:"config,omitempty"`
		General struct {
			ExpertMode     string `json:"expertMode,omitempty"`
			AdminPassword  string `json:"admin-password,omitempty"`
			Nolog          string `json:"nolog,omitempty"`
			RconPassword   string `json:"rcon-password,omitempty"`
			AdditionalMods string `json:"additionalMods,omitempty"`
			Bans           string `json:"bans,omitempty"`
			Whitelist      string `json:"whitelist,omitempty"`
			Resetmission   string `json:"resetmission,omitempty"`
			Priority       string `json:"priority,omitempty"`
		} `json:"general,omitempty"`
	} `json:"settings,omitempty"`
	Quota interface{} `json:"quota,omitempty"`
	Query struct {
		ServerName    string `json:"server_name,omitempty"`
		ConnectIP     string `json:"connect_ip,omitempty"`
		Map           string `json:"map,omitempty"`
		Version       string `json:"version,omitempty"`
		PlayerCurrent int    `json:"player_current,omitempty"`
		PlayerMax     int    `json:"player_max,omitempty"`
	} `json:"query,omitempty"`
}

// GameServerDetailResp contains the query response from the Nitrado API for the GameServer operation
type GameServerDetailResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		GameServer GameServer `json:"gameserver,omitempty"`
	} `json:"data,omitempty"`
}

// GameServerRestartResp contains the query response from the Nitrado API for the GameServer operation
type GameServerRestartResp struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// Get a GameServer by service ID.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-Details
func (s *GameServersService) Get(serviceID int) (*GameServer, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers", serviceID)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gameServerDetailResp *GameServerDetailResp
	resp, err := s.client.Do(req, &gameServerDetailResp)
	if err != nil {
		return nil, resp, err
	}

	return &gameServerDetailResp.Data.GameServer, resp, nil
}

// Restart a GameServer by service ID.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-Details
func (s *GameServersService) Restart(serviceID int) error {
	u := fmt.Sprintf("services/%v/gameservers/restart", serviceID)
	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return err
	}

	var gameServerRestartResp *GameServerRestartResp
	_, err = s.client.Do(req, &gameServerRestartResp)
	if err != nil {
		return err
	}
	if gameServerRestartResp.Status != "success" {
		return fmt.Errorf("status %q (%q)", gameServerRestartResp.Status, gameServerRestartResp.Message)
	}

	return nil
}

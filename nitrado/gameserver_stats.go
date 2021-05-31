package nitrado

import (
	"fmt"
	"net/http"
)

// Generated structs from https://mholt.github.io/json-to-go/

// GameServerStatsService provides access to the service related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/
type GameServerStatsService apiService

// GSStatsResp contains the response object from the stats method on a gameserver
type GSStatsResp struct {
	Status string `json:"status"`
	Data   struct {
		Stats GSStats `json:"stats"`
	} `json:"data"`
}

// GSStats contains a stats object for a game server
type GSStats struct {
	CPUUsage       [][]float32 `json:"cpuUsage"`
	CurrentPlayers [][]float32 `json:"currentPlayers"`
	MaxPlayers     [][]float32 `json:"maxPlayers"`
	MemoryUsage    [][]float32 `json:"memoryUsage"`
}

// Get stats from a GameServer by service ID.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-Stats
func (s *GameServerStatsService) Get(serviceID int) (*GSStats, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/stats", serviceID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var gameServerStatsResp *GSStatsResp
	resp, err := s.client.Do(req, &gameServerStatsResp)
	if err != nil {
		return nil, resp, err
	}

	return &gameServerStatsResp.Data.Stats, resp, nil
}

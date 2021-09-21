package nitrado

import (
	"fmt"
	"net/http"
	"sort"
)

// Generated structs from https://mholt.github.io/json-to-go/

// PlayerListService provides access to the player list related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Player_Management-ListPlayers
type PlayerListService apiService

// Player contains the details of a player
type Player struct {
	Name    string   `json:"name,omitempty"`
	ID      string   `json:"id,omitempty"`
	IDType  string   `json:"id_type,omitempty"`
	Online  string   `json:"online,omitempty"`
	Actions []string `json:"actions,omitempty"`
}

// PlayerListResp contains a listing of the players for a gameserver
type PlayerListResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Players []Player `json:"players,omitempty"`
	} `json:"data,omitempty"`
}

// List players on a GameServer.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-GameserverFilesList
func (s *PlayerListService) List(svc Service) ([]Player, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/games/players", svc.ID)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var playerListResp *PlayerListResp
	resp, err := s.client.Do(req, &playerListResp)
	if err != nil {
		return nil, resp, err
	}

	// Sort the files by Modified date
	sort.Slice(playerListResp.Data.Players, func(i, j int) bool {
		return playerListResp.Data.Players[i].Name < playerListResp.Data.Players[j].Name
	})

	return playerListResp.Data.Players, resp, nil
}

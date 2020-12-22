package nitrado

import (
	"fmt"
	"net/http"
	"sort"
)

// Generated structs from https://mholt.github.io/json-to-go/

// FileServerService provides access to the service related functions in the Nitrado API.
//
// Nitrado API docs: https://doc.nitrado.net/
type FileServerService apiService

// File contains the details of a file
type File struct {
	Owner      string `json:"owner,omitempty"`
	CreatedAt  int    `json:"created_at,omitempty"`
	Path       string `json:"path,omitempty"`
	Size       int    `json:"size,omitempty"`
	AccessedAt int    `json:"accessed_at,omitempty"`
	ModifiedAt int    `json:"modified_at,omitempty"`
	Type       string `json:"type,omitempty"`
	Chmod      string `json:"chmod,omitempty"`
	Group      string `json:"group,omitempty"`
	Name       string `json:"name,omitempty"`
}

// FileListResp contains a listing of the files at a location
type FileListResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Entries []File `json:"entries,omitempty"`
	} `json:"data,omitempty"`
}

// List files on a GameServer.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-GameserverFilesList
func (s *FileServerService) List(svc Service, dir string) (*[]File, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/file_server/list", svc.ID)
	if dir != "" {
		u = fmt.Sprintf("%s?dir=%s", u, dir)
	}
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var fileListResp *FileListResp
	resp, err := s.client.Do(req, &fileListResp)
	if err != nil {
		return nil, resp, err
	}

	// Sort the files by Modified date
	sort.Slice(fileListResp.Data.Entries, func(i, j int) bool {
		return fileListResp.Data.Entries[i].ModifiedAt < fileListResp.Data.Entries[j].ModifiedAt
	})

	return &fileListResp.Data.Entries, resp, nil
}

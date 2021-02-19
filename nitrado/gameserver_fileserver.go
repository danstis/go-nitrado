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

// FileDownloadResp contains the response object from the download method on a fileserver
type FileDownloadResp struct {
	Status string `json:"status,omitempty"`
	Data   struct {
		Token struct {
			URL   string `json:"url,omitempty"`
			Token string `json:"token,omitempty"`
		} `json:"token,omitempty"`
	} `json:"data,omitempty"`
}

// FileServerListOptions controls the query string settings that a list request can take.
type FileServerListOptions struct {
	Dir string `url:"dir,omitempty"`
}

// FileServerDownloadOptions controls the query string settings that a download request can take.
type FileServerDownloadOptions struct {
	File string `url:"file,omitempty"`
}

// FileServerUploadOptions controls the query string settings that an upload request can take.
type FileServerUploadOptions struct {
	Path string `url:"path,omitempty"`
	File string `url:"file,omitempty"`
}

// List files on a GameServer.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-GameserverFilesList
func (s *FileServerService) List(svc Service, opts FileServerListOptions) ([]File, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/file_server/list", svc.ID)
	u, err := addOptions(u, opts)
	if err != nil {
		return nil, nil, err
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

	return fileListResp.Data.Entries, resp, nil
}

// Download a given file on a GameServer.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-GameserverFilesDownload
func (s *FileServerService) Download(svc Service, opts FileServerDownloadOptions) (string, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/file_server/download", svc.ID)
	u, err := addOptions(u, opts)
	if err != nil {
		return "", nil, err
	}

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return "", nil, err
	}

	var fileDownloadResp *FileDownloadResp
	resp, err := s.client.Do(req, &fileDownloadResp)
	if err != nil {
		return "", resp, err
	}

	return fileDownloadResp.Data.Token.URL, resp, nil
}

// Upload a given file on a GameServer.
//
// Nitrado API docs: https://doc.nitrado.net/#api-Gameserver-GameserverFilesUpload
func (s *FileServerService) Upload(svc Service, opts FileServerUploadOptions) (FileDownloadResp, *http.Response, error) {
	u := fmt.Sprintf("services/%v/gameservers/file_server/upload", svc.ID)
	u, err := addOptions(u, opts)
	if err != nil {
		return FileDownloadResp{}, nil, err
	}

	req, err := s.client.NewRequest("POST", u, nil)
	if err != nil {
		return FileDownloadResp{}, nil, err
	}

	var fileDownloadResp *FileDownloadResp
	resp, err := s.client.Do(req, &fileDownloadResp)
	if err != nil {
		return FileDownloadResp{}, resp, err
	}

	return *fileDownloadResp, resp, nil
}

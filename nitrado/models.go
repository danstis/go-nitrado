package nitrado

// Generated structs from https://mholt.github.io/json-to-go/

// FileBookmarks contains the list of file locations on a game server
type FileBookmarks struct {
	Status string `json:"status"`
	Data   struct {
		Bookmarks []string `json:"bookmarks"`
	} `json:"data"`
}

// File contains a list of files
type File struct {
	Owner      string `json:"owner"`
	CreatedAt  int    `json:"created_at"`
	Path       string `json:"path"`
	Size       int    `json:"size"`
	AccessedAt int    `json:"accessed_at"`
	ModifiedAt int    `json:"modified_at"`
	Type       string `json:"type"`
	Chmod      string `json:"chmod"`
	Group      string `json:"group"`
	Name       string `json:"name"`
}

// FileList contains a listing of the files at a location
type FileList struct {
	Status string `json:"status"`
	Data   struct {
		Entries []File `json:"entries"`
	} `json:"data"`
}

// FileLink contains a download link to a file
type FileLink struct {
	Status string `json:"status"`
	Data   struct {
		Token struct {
			URL   string `json:"url"`
			Token string `json:"token"`
		} `json:"token"`
	} `json:"data"`
}

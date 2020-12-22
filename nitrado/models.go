package nitrado

// Generated structs from https://mholt.github.io/json-to-go/

// FileBookmarks contains the list of file locations on a game server
type FileBookmarks struct {
	Status string `json:"status"`
	Data   struct {
		Bookmarks []string `json:"bookmarks"`
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

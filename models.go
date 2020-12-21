package nitrado

// Generated structs from https://mholt.github.io/json-to-go/

// Service contains the structure of a nitrado.net Api service object
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

// Services contains a list of services
type Services struct {
	Status string `json:"status"`
	Data   struct {
		Services []Service `json:"services"`
	} `json:"data"`
}

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

// GameServer contains a Game Server details
type GameServer struct {
	Status           string `json:"status"`
	LastStatusChange int    `json:"last_status_change"`
	MustBeStarted    bool   `json:"must_be_started"`
	WebsocketToken   string `json:"websocket_token"`
	Hostsystems      struct {
		Linux struct {
			Hostname   string `json:"hostname"`
			Servername string `json:"servername"`
			Status     string `json:"status"`
		} `json:"linux"`
		Windows struct {
			Hostname   string `json:"hostname"`
			Servername string `json:"servername"`
			Status     string `json:"status"`
		} `json:"windows"`
	} `json:"hostsystems"`
	Username      string      `json:"username"`
	UserID        int         `json:"user_id"`
	ServiceID     int         `json:"service_id"`
	LocationID    int         `json:"location_id"`
	MinecraftMode bool        `json:"minecraft_mode"`
	IP            string      `json:"ip"`
	Ipv6          interface{} `json:"ipv6"`
	Port          int         `json:"port"`
	QueryPort     int         `json:"query_port"`
	RconPort      int         `json:"rcon_port"`
	Label         string      `json:"label"`
	Type          string      `json:"type"`
	Memory        string      `json:"memory"`
	MemoryMb      int         `json:"memory_mb"`
	Game          string      `json:"game"`
	GameHuman     string      `json:"game_human"`
	GameSpecific  struct {
		Path          string `json:"path"`
		UpdateStatus  string `json:"update_status"`
		LastUpdate    string `json:"last_update"`
		PathAvailable bool   `json:"path_available"`
		Features      struct {
			HasBackups                 bool `json:"has_backups"`
			HasWorldBackups            bool `json:"has_world_backups"`
			HasRcon                    bool `json:"has_rcon"`
			HasApplicationServer       bool `json:"has_application_server"`
			HasContainerWebsocket      bool `json:"has_container_websocket"`
			HasFileBrowser             bool `json:"has_file_browser"`
			HasFtp                     bool `json:"has_ftp"`
			HasExpertMode              bool `json:"has_expert_mode"`
			HasPackages                bool `json:"has_packages"`
			HasPluginSystem            bool `json:"has_plugin_system"`
			HasRestartMessageSupport   bool `json:"has_restart_message_support"`
			HasDatabase                bool `json:"has_database"`
			HasPlayermanagementFeature bool `json:"has_playermanagement_feature"`
		} `json:"features"`
		LogFiles    []string      `json:"log_files"`
		ConfigFiles []interface{} `json:"config_files"`
	} `json:"game_specific"`
	Slots       int    `json:"slots"`
	Location    string `json:"location"`
	Credentials struct {
		Ftp struct {
			Hostname string `json:"hostname"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"ftp"`
		Mysql struct {
			Hostname string `json:"hostname"`
			Port     int    `json:"port"`
			Username string `json:"username"`
			Password string `json:"password"`
			Database string `json:"database"`
		} `json:"mysql"`
	} `json:"credentials"`
	Settings struct {
		Config struct {
			Hostname                    string `json:"hostname"`
			VonCodecQuality             string `json:"vonCodecQuality"`
			DisableVoN                  string `json:"disableVoN"`
			Password                    string `json:"password"`
			ServerTimeAcceleration      string `json:"serverTimeAcceleration"`
			ServerNightTimeAcceleration string `json:"serverNightTimeAcceleration"`
			ServerTimePersistent        string `json:"serverTimePersistent"`
			Disable3RdPerson            string `json:"disable3rdPerson"`
			DisableCrosshair            string `json:"disableCrosshair"`
			UseServerTime               string `json:"useServerTime"`
			CustomServerTime            string `json:"customServerTime"`
			EnableMouseAndKeyboard      string `json:"enableMouseAndKeyboard"`
			EnableWhitelist             string `json:"enableWhitelist"`
			Mission                     string `json:"mission"`
			AdminLogPlayerHitsOnly      string `json:"adminLogPlayerHitsOnly"`
			AdminLogPlacement           string `json:"adminLogPlacement"`
			AdminLogBuildActions        string `json:"adminLogBuildActions"`
			AdminLogPlayerList          string `json:"adminLogPlayerList"`
			LightingConfig              string `json:"lightingConfig"`
			DisablePersonalLight        string `json:"disablePersonalLight"`
			DisableBaseDamage           string `json:"disableBaseDamage"`
			DisableContainerDamage      string `json:"disableContainerDamage"`
		} `json:"config"`
		General struct {
			ExpertMode     string `json:"expertMode"`
			AdminPassword  string `json:"admin-password"`
			Nolog          string `json:"nolog"`
			RconPassword   string `json:"rcon-password"`
			AdditionalMods string `json:"additionalMods"`
			Bans           string `json:"bans"`
			Whitelist      string `json:"whitelist"`
			Resetmission   string `json:"resetmission"`
			Priority       string `json:"priority"`
		} `json:"general"`
	} `json:"settings"`
	Quota interface{} `json:"quota"`
	Query struct {
		ServerName    string        `json:"server_name"`
		ConnectIP     string        `json:"connect_ip"`
		Map           string        `json:"map"`
		Version       string        `json:"version"`
		PlayerCurrent int           `json:"player_current"`
		PlayerMax     int           `json:"player_max"`
		Players       []interface{} `json:"players"`
	} `json:"query"`
}

// GameServerQuery contains the query response from the Nitrado API for the GameServer operation
type GameServerQuery struct {
	Status string `json:"status"`
	Data   struct {
		GameServer GameServer `json:"gameserver"`
	} `json:"data"`
}

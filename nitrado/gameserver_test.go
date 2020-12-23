package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGameServersService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"success","data":{"gameserver":{"status":"started","last_status_change":1608612373,"must_be_started":true,"websocket_token":"abcdefgh012345","hostsystems":{"linux":{"hostname":"ausy001.nitrado.net","servername":"ausy001","status":"online"},"windows":{"hostname":"ausy002.nitrado.net","servername":"ausy002","status":"online"}},"username":"ni2_2","user_id":1234567,"service_id":7654321,"location_id":14,"minecraft_mode":false,"ip":"128.0.0.123","ipv6":null,"port":10900,"query_port":10903,"rcon_port":10903,"label":"ni","type":"Gameserver","memory":"Standard","memory_mb":4096,"game":"dayzxb","game_human":"DayZ (Xbox One)","game_specific":{"path":"/games/ni2_2/noftp/dayzxb/","update_status":"up_to_date","last_update":"2020-12-01T01:11:08","path_available":true,"features":{"has_backups":true,"has_world_backups":false,"has_rcon":false,"has_application_server":false,"has_container_websocket":false,"has_file_browser":true,"has_ftp":true,"has_expert_mode":false,"has_packages":false,"has_plugin_system":false,"has_restart_message_support":false,"has_database":true,"has_playermanagement_feature":false},"log_files":["dayzxb/config/DayZServer_X1_x64.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_22_054608878.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_22_023238448.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_21_231926325.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_21_200630133.ADM"],"config_files":[]},"modpacks":{},"slots":50,"location":"AU","credentials":{"ftp":{"hostname":"ausy001.gamedata.io","port":21,"username":"ni2_2","password":"ABC123z"},"mysql":{"hostname":"ausy001.gamedata.io","port":3306,"username":"ni2_2_DB","password":"ABC123z","database":"ni2_2_DB"}},"settings":{"config":{"hostname":"Server Name","vonCodecQuality":"20","disableVoN":"0","password":"","serverTimeAcceleration":"4.8","serverNightTimeAcceleration":"5","serverTimePersistent":"0","disable3rdPerson":"0","disableCrosshair":"0","useServerTime":"0","customServerTime":"2020/07/01/04/00","enableMouseAndKeyboard":"1","enableWhitelist":"0","mission":"dayzOffline.chernarusplus","adminLogPlayerHitsOnly":"1","adminLogPlacement":"1","adminLogBuildActions":"1","adminLogPlayerList":"1","lightingConfig":"1","disablePersonalLight":"0","disableBaseDamage":"0","disableContainerDamage":"1","disableRespawnDialog":"0"},"general":{"expertMode":"false","admin-password":"ABC_123z","nolog":"false","rcon-password":"","additionalMods":"","bans":"Gamer Tag 1\r\nGamer Tag 2","whitelist":"Gamer Tag 1\r\nGamer Tag 2","resetmission":"false","priority":"Gamer Tag 1\r\nGamer Tag 2"}},"quota":null,"query":{"server_name":"Server Name","connect_ip":"128.0.0.123:10900","map":"dayzOffline.chernarusplus","version":"v1.10.153598","player_current":23,"player_max":50,"players":[]}}}}`)
	})

	type args struct {
		serviceID int
	}
	tests := []struct {
		name    string
		s       *GameServersService
		args    args
		want    *GameServer
		wantErr bool
	}{
		{
			name: "List Services",
			s:    client.GameServers,
			args: args{serviceID: 7654321},
			want: &GameServer{
				Status:           "started",
				LastStatusChange: 1608612373,
				MustBeStarted:    true,
				Username:         "ni2_2",
				UserID:           1234567,
				ServiceID:        7654321,
				IP:               "128.0.0.123",
				Port:             10900,
				QueryPort:        10903,
				RconPort:         10903,
				Type:             "Gameserver",
				Memory:           "Standard",
				MemoryMb:         4096,
				Game:             "dayzxb",
				GameHuman:        "DayZ (Xbox One)",
				GameSpecific: struct {
					Path         string   `json:"path,omitempty"`
					UpdateStatus string   `json:"update_status,omitempty"`
					LastUpdate   string   `json:"last_update,omitempty"`
					LogFiles     []string `json:"log_files,omitempty"`
					ConfigFiles  []string `json:"config_files,omitempty"`
				}{
					Path:         "/games/ni2_2/noftp/dayzxb/",
					UpdateStatus: "up_to_date",
					LastUpdate:   "2020-12-01T01:11:08",
					LogFiles: []string{
						"dayzxb/config/DayZServer_X1_x64.ADM",
						"dayzxb/config/DayZServer_X1_x64_2020_12_22_054608878.ADM",
						"dayzxb/config/DayZServer_X1_x64_2020_12_22_023238448.ADM",
						"dayzxb/config/DayZServer_X1_x64_2020_12_21_231926325.ADM",
						"dayzxb/config/DayZServer_X1_x64_2020_12_21_200630133.ADM",
					},
					ConfigFiles: []string{},
				},
				Slots:    50,
				Location: "AU",
				Credentials: struct {
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
				}{
					Ftp: struct {
						Hostname string `json:"hostname,omitempty"`
						Port     int    `json:"port,omitempty"`
						Username string `json:"username,omitempty"`
						Password string `json:"password,omitempty"`
					}{
						Hostname: "ausy001.gamedata.io",
						Port:     21,
						Username: "ni2_2",
						Password: "ABC123z",
					},
					Mysql: struct {
						Hostname string `json:"hostname,omitempty"`
						Port     int    `json:"port,omitempty"`
						Username string `json:"username,omitempty"`
						Password string `json:"password,omitempty"`
						Database string `json:"database,omitempty"`
					}{
						Hostname: "ausy001.gamedata.io",
						Port:     3306,
						Username: "ni2_2_DB",
						Password: "ABC123z",
						Database: "ni2_2_DB",
					},
				},
				Settings: struct {
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
				}{
					Config: struct {
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
					}{
						Hostname:                    "Server Name",
						VonCodecQuality:             "20",
						DisableVoN:                  "0",
						Password:                    "",
						ServerTimeAcceleration:      "4.8",
						ServerNightTimeAcceleration: "5",
						ServerTimePersistent:        "0",
						Disable3RdPerson:            "0",
						DisableCrosshair:            "0",
						UseServerTime:               "0",
						CustomServerTime:            "2020/07/01/04/00",
						EnableMouseAndKeyboard:      "1",
						EnableWhitelist:             "0",
						Mission:                     "dayzOffline.chernarusplus",
						AdminLogPlayerHitsOnly:      "1",
						AdminLogPlacement:           "1",
						AdminLogBuildActions:        "1",
						AdminLogPlayerList:          "1",
						LightingConfig:              "1",
						DisablePersonalLight:        "0",
						DisableBaseDamage:           "0",
						DisableContainerDamage:      "1",
					},
					General: struct {
						ExpertMode     string `json:"expertMode,omitempty"`
						AdminPassword  string `json:"admin-password,omitempty"`
						Nolog          string `json:"nolog,omitempty"`
						RconPassword   string `json:"rcon-password,omitempty"`
						AdditionalMods string `json:"additionalMods,omitempty"`
						Bans           string `json:"bans,omitempty"`
						Whitelist      string `json:"whitelist,omitempty"`
						Resetmission   string `json:"resetmission,omitempty"`
						Priority       string `json:"priority,omitempty"`
					}{
						ExpertMode:     "false",
						AdminPassword:  "ABC_123z",
						Nolog:          "false",
						RconPassword:   "",
						AdditionalMods: "",
						Bans:           "Gamer Tag 1\r\nGamer Tag 2",
						Whitelist:      "Gamer Tag 1\r\nGamer Tag 2",
						Resetmission:   "false",
						Priority:       "Gamer Tag 1\r\nGamer Tag 2",
					},
				},
				Quota: nil,
				Query: struct {
					ServerName    string `json:"server_name,omitempty"`
					ConnectIP     string `json:"connect_ip,omitempty"`
					Map           string `json:"map,omitempty"`
					Version       string `json:"version,omitempty"`
					PlayerCurrent int    `json:"player_current,omitempty"`
					PlayerMax     int    `json:"player_max,omitempty"`
				}{
					ServerName:    "Server Name",
					ConnectIP:     "128.0.0.123:10900",
					Map:           "dayzOffline.chernarusplus",
					Version:       "v1.10.153598",
					PlayerCurrent: 23,
					PlayerMax:     50,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.Get(tt.args.serviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameServersService.Get() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameServersService.Get() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestGameServersService_Restart(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/restart", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"status":"success","message":"Server will be restarted now."}`)
	})
	mux.HandleFunc("/services/999/gameservers/restart", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"status":"failure","message":"Test failure."}`)
	})

	type args struct {
		serviceID int
	}
	tests := []struct {
		name    string
		s       *GameServersService
		args    args
		wantErr bool
	}{
		{
			name:    "Restart gameserver",
			s:       client.GameServers,
			args:    args{serviceID: 7654321},
			wantErr: false,
		},
		{
			name:    "Restart gameserver",
			s:       client.GameServers,
			args:    args{serviceID: 999},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Restart(tt.args.serviceID); (err != nil) != tt.wantErr {
				t.Errorf("GameServersService.Restart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

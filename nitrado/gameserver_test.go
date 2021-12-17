package nitrado

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// JSON minified using https://codebeautify.org/jsonminifier

//TestGameServersService_Get tests the GameServersService Get() method.
func TestGameServersService_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"success","data":{"gameserver":{"status":"started","last_status_change":1608612373,"must_be_started":true,"websocket_token":"abcdefgh012345","hostsystems":{"linux":{"hostname":"ausy001.nitrado.net","servername":"ausy001","status":"online"},"windows":{"hostname":"ausy002.nitrado.net","servername":"ausy002","status":"online"}},"username":"ni2_2","user_id":1234567,"service_id":7654321,"location_id":14,"minecraft_mode":false,"ip":"128.0.0.123","ipv6":null,"port":10900,"query_port":10903,"rcon_port":10903,"label":"ni","type":"Gameserver","memory":"Standard","memory_mb":4096,"game":"dayzxb","game_human":"DayZ (Xbox One)","game_specific":{"path":"/games/ni2_2/noftp/dayzxb/","update_status":"up_to_date","last_update":"2020-12-01T01:11:08","path_available":true,"features":{"has_backups":true,"has_world_backups":false,"has_rcon":false,"has_application_server":false,"has_container_websocket":false,"has_file_browser":true,"has_ftp":true,"has_expert_mode":false,"has_packages":false,"has_plugin_system":false,"has_restart_message_support":false,"has_database":true,"has_playermanagement_feature":false},"log_files":["dayzxb/config/DayZServer_X1_x64.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_22_054608878.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_22_023238448.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_21_231926325.ADM","dayzxb/config/DayZServer_X1_x64_2020_12_21_200630133.ADM"],"config_files":[]},"modpacks":{},"slots":50,"location":"AU","credentials":{"ftp":{"hostname":"ausy001.gamedata.io","port":21,"username":"ni2_2","password":"ABC123z"},"mysql":{"hostname":"ausy001.gamedata.io","port":3306,"username":"ni2_2_DB","password":"ABC123z","database":"ni2_2_DB"}},"settings":{"config":{"hostname":"Server Name","vonCodecQuality":"20","disableVoN":"0","password":"","serverTimeAcceleration":"4.8","serverNightTimeAcceleration":"5","serverTimePersistent":"0","disable3rdPerson":"0","disableCrosshair":"0","useServerTime":"0","customServerTime":"2020/07/01/04/00","enableMouseAndKeyboard":"1","enableWhitelist":"0","mission":"dayzOffline.chernarusplus","adminLogPlayerHitsOnly":"1","adminLogPlacement":"1","adminLogBuildActions":"1","adminLogPlayerList":"1","lightingConfig":"1","disablePersonalLight":"0","disableBaseDamage":"0","disableContainerDamage":"1","disableRespawnDialog":"0","enableCfgGameplayFile":"1"},"general":{"expertMode":"false","admin-password":"ABC_123z","nolog":"false","rcon-password":"","additionalMods":"","bans":"Gamer Tag 1\r\nGamer Tag 2","whitelist":"Gamer Tag 1\r\nGamer Tag 2","resetmission":"false","priority":"Gamer Tag 1\r\nGamer Tag 2"}},"quota":null,"query":{"server_name":"Server Name","connect_ip":"128.0.0.123:10900","map":"dayzOffline.chernarusplus","version":"v1.10.153598","player_current":23,"player_max":50,"players":[]}}}}`)
	})

	s := client.GameServers
	got, _, err := s.Get(7654321)

	require.Nil(t, err)

	assert.Equal(t, "started", got.Status)
	assert.Equal(t, int(1608612373), got.LastStatusChange)
	assert.Equal(t, true, got.MustBeStarted)
	assert.Equal(t, "ni2_2", got.Username)
	assert.Equal(t, int(1234567), got.UserID)
	assert.Equal(t, int(7654321), got.ServiceID)
	assert.Equal(t, "128.0.0.123", got.IP)
	assert.Equal(t, int(10900), got.Port)
	assert.Equal(t, int(10903), got.QueryPort)
	assert.Equal(t, int(10903), got.RconPort)
	assert.Equal(t, "Gameserver", got.Type)
	assert.Equal(t, "Standard", got.Memory)
	assert.Equal(t, int(4096), got.MemoryMb)
	assert.Equal(t, "dayzxb", got.Game)
	assert.Equal(t, "DayZ (Xbox One)", got.GameHuman)
	assert.Equal(t, 50, got.Slots)
	assert.Equal(t, "AU", got.Location)
	assert.Equal(t, nil, got.Quota)

	// GameSpecific
	assert.Equal(t, "/games/ni2_2/noftp/dayzxb/", got.GameSpecific.Path)
	assert.Equal(t, "up_to_date", got.GameSpecific.UpdateStatus)
	assert.Equal(t, "2020-12-01T01:11:08", got.GameSpecific.LastUpdate)
	assert.Equal(t, 5, len(got.GameSpecific.LogFiles))
	assert.Contains(t, got.GameSpecific.LogFiles, "dayzxb/config/DayZServer_X1_x64.ADM")
	assert.Contains(t, got.GameSpecific.LogFiles, "dayzxb/config/DayZServer_X1_x64_2020_12_22_054608878.ADM")
	assert.Contains(t, got.GameSpecific.LogFiles, "dayzxb/config/DayZServer_X1_x64_2020_12_22_023238448.ADM")
	assert.Contains(t, got.GameSpecific.LogFiles, "dayzxb/config/DayZServer_X1_x64_2020_12_21_231926325.ADM")
	assert.Contains(t, got.GameSpecific.LogFiles, "dayzxb/config/DayZServer_X1_x64_2020_12_21_200630133.ADM")

	// Credentials
	assert.Equal(t, "ausy001.gamedata.io", got.Credentials.Ftp.Hostname)
	assert.Equal(t, int(21), got.Credentials.Ftp.Port)
	assert.Equal(t, "ni2_2", got.Credentials.Ftp.Username)
	assert.Equal(t, "ABC123z", got.Credentials.Ftp.Password)
	assert.Equal(t, "ausy001.gamedata.io", got.Credentials.Mysql.Hostname)
	assert.Equal(t, int(3306), got.Credentials.Mysql.Port)
	assert.Equal(t, "ni2_2_DB", got.Credentials.Mysql.Username)
	assert.Equal(t, "ABC123z", got.Credentials.Mysql.Password)
	assert.Equal(t, "ni2_2_DB", got.Credentials.Mysql.Database)
	// Settings.Config
	assert.Equal(t, "Server Name", got.Settings.Config.Hostname)
	assert.Equal(t, "20", got.Settings.Config.VonCodecQuality)
	assert.Equal(t, "0", got.Settings.Config.DisableVoN)
	assert.Equal(t, "", got.Settings.Config.Password)
	assert.Equal(t, "4.8", got.Settings.Config.ServerTimeAcceleration)
	assert.Equal(t, "5", got.Settings.Config.ServerNightTimeAcceleration)
	assert.Equal(t, "0", got.Settings.Config.ServerTimePersistent)
	assert.Equal(t, "0", got.Settings.Config.Disable3RdPerson)
	assert.Equal(t, "0", got.Settings.Config.DisableCrosshair)
	assert.Equal(t, "0", got.Settings.Config.UseServerTime)
	assert.Equal(t, "2020/07/01/04/00", got.Settings.Config.CustomServerTime)
	assert.Equal(t, "1", got.Settings.Config.EnableMouseAndKeyboard)
	assert.Equal(t, "0", got.Settings.Config.EnableWhitelist)
	assert.Equal(t, "dayzOffline.chernarusplus", got.Settings.Config.Mission)
	assert.Equal(t, "1", got.Settings.Config.AdminLogPlayerHitsOnly)
	assert.Equal(t, "1", got.Settings.Config.AdminLogPlacement)
	assert.Equal(t, "1", got.Settings.Config.AdminLogBuildActions)
	assert.Equal(t, "1", got.Settings.Config.AdminLogPlayerList)
	assert.Equal(t, "1", got.Settings.Config.LightingConfig)
	assert.Equal(t, "0", got.Settings.Config.DisablePersonalLight)
	assert.Equal(t, "0", got.Settings.Config.DisableBaseDamage)
	assert.Equal(t, "1", got.Settings.Config.DisableContainerDamage)
	assert.Equal(t, "1", got.Settings.Config.EnableCFGGameplayFile)
	// Settings.General
	assert.Equal(t, "false", got.Settings.General.ExpertMode)
	assert.Equal(t, "ABC_123z", got.Settings.General.AdminPassword)
	assert.Equal(t, "false", got.Settings.General.Nolog)
	assert.Equal(t, "", got.Settings.General.RconPassword)
	assert.Equal(t, "", got.Settings.General.AdditionalMods)
	assert.Equal(t, "Gamer Tag 1\r\nGamer Tag 2", got.Settings.General.Bans)
	assert.Equal(t, "Gamer Tag 1\r\nGamer Tag 2", got.Settings.General.Whitelist)
	assert.Equal(t, "false", got.Settings.General.Resetmission)
	assert.Equal(t, "Gamer Tag 1\r\nGamer Tag 2", got.Settings.General.Priority)

	// Query
	assert.Equal(t, "Server Name", got.Query.ServerName)
	assert.Equal(t, "128.0.0.123:10900", got.Query.ConnectIP)
	assert.Equal(t, "dayzOffline.chernarusplus", got.Query.Map)
	assert.Equal(t, "v1.10.153598", got.Query.Version)
	assert.Equal(t, 23, got.Query.PlayerCurrent)
	assert.Equal(t, 50, got.Query.PlayerMax)
}

// TestGameServersService_Restart tests the GameServersService Restart() method.
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

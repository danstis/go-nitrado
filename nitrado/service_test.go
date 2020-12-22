package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServicesService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"success","data":{"services":[{"id":3,"location_id":2,"status":"active","websocket_token":"abcdefgh012345","user_id":2,"comment":"This is my special Battlefield Server.","auto_extension":false,"auto_extension_duration":null,"type":"gameserver","type_human":"Publicserver 16 Slots","managedroot_id":null,"details":{"address":"10.10.0.6:27015","name":"Nitrado.net Battlefield 4 Server","game":"Battlefield 4","portlist_short":"bf4","folder_short":"bf4","slots":16},"start_date":"2015-08-11T13:01:01","suspend_date":"2017-02-09T22:26:46","delete_date":"2017-02-19T22:26:46","username":"ni2_1","roles":["ROLE_OWNER"]},{"id":6,"location_id":2,"status":"active","websocket_token":"abcdefgh012345","user_id":1,"comment":null,"auto_extension":false,"auto_extension_duration":null,"type":"gameserver","type_human":"Publicserver 4 Slots","managedroot_id":826777,"details":{"address":"10.10.0.7:27015","name":"Nitrado.net Minecraft Server","game":"Minecraft","portlist_short":"mcr","folder_short":"minecraft","slots":4},"start_date":"2015-08-11T13:01:01","suspend_date":"2017-02-09T22:26:46","delete_date":"2017-02-19T22:26:46","username":"ni1_1","roles":["ROLE_GAMESERVER_CHANGE_GAME","ROLE_WEBINTERFACE_GENERAL_CONTROL","ROLE_WEBINTERFACE_SETTINGS_READ","ROLE_WEBINTERFACE_SETTINGS_WRITE","ROLE_WEBINTERFACE_LOGS_READ","ROLE_WEBINTERFACE_SCHEDULED_RESTART_READ","ROLE_WEBINTERFACE_SCHEDULED_RESTART_WRITE"]}]}}`)
	})

	tests := []struct {
		name    string
		s       *ServicesService
		want    *[]Service
		wantErr bool
	}{
		{
			name: "List Services",
			s:    client.Services,
			want: &[]Service{
				{
					ID:                    3,
					LocationID:            2,
					Status:                "active",
					WebsocketToken:        "abcdefgh012345",
					UserID:                2,
					Comment:               "This is my special Battlefield Server.",
					AutoExtension:         false,
					AutoExtensionDuration: 0,
					Type:                  "gameserver",
					TypeHuman:             "Publicserver 16 Slots",
					Details: struct {
						Address       string "json:\"address\""
						Name          string "json:\"name\""
						Game          string "json:\"game\""
						PortlistShort string "json:\"portlist_short\""
						FolderShort   string "json:\"folder_short\""
						Slots         int    "json:\"slots\""
					}{
						Address:       "10.10.0.6:27015",
						Name:          "Nitrado.net Battlefield 4 Server",
						Game:          "Battlefield 4",
						PortlistShort: "bf4",
						FolderShort:   "bf4",
						Slots:         16,
					},
					StartDate:   "2015-08-11T13:01:01",
					SuspendDate: "2017-02-09T22:26:46",
					DeleteDate:  "2017-02-19T22:26:46",
					Username:    "ni2_1",
					Roles:       []string{"ROLE_OWNER"},
				},
				{
					ID:                    6,
					LocationID:            2,
					Status:                "active",
					WebsocketToken:        "abcdefgh012345",
					UserID:                1,
					Comment:               nil,
					AutoExtension:         false,
					AutoExtensionDuration: 0,
					Type:                  "gameserver",
					TypeHuman:             "Publicserver 4 Slots",
					Details: struct {
						Address       string "json:\"address\""
						Name          string "json:\"name\""
						Game          string "json:\"game\""
						PortlistShort string "json:\"portlist_short\""
						FolderShort   string "json:\"folder_short\""
						Slots         int    "json:\"slots\""
					}{
						Address:       "10.10.0.7:27015",
						Name:          "Nitrado.net Minecraft Server",
						Game:          "Minecraft",
						PortlistShort: "mcr",
						FolderShort:   "minecraft",
						Slots:         4,
					},
					StartDate:   "2015-08-11T13:01:01",
					SuspendDate: "2017-02-09T22:26:46",
					DeleteDate:  "2017-02-19T22:26:46",
					Username:    "ni1_1",
					Roles: []string{
						"ROLE_GAMESERVER_CHANGE_GAME",
						"ROLE_WEBINTERFACE_GENERAL_CONTROL",
						"ROLE_WEBINTERFACE_SETTINGS_READ",
						"ROLE_WEBINTERFACE_SETTINGS_WRITE",
						"ROLE_WEBINTERFACE_LOGS_READ",
						"ROLE_WEBINTERFACE_SCHEDULED_RESTART_READ",
						"ROLE_WEBINTERFACE_SCHEDULED_RESTART_WRITE",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("ServicesService.List() error = %#v, wantErr %#v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ServicesService.List() got = %#v, want %#v", got, tt.want)
			}
		})
	}
}

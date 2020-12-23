package nitrado

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
)

func TestGSSettingsService_Update(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"status":"success","data":{"settings":{"config":{"mysetting":"true"}}}}`)
	})
	mux.HandleFunc("/services/999/gameservers/settings", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		fmt.Fprint(w, `{"status":"failure"}`)
	})

	type args struct {
		serviceID int
		opts      GSSettingsUpdateOptions
	}
	tests := []struct {
		name    string
		s       *GSSettingsService
		args    args
		wantErr bool
	}{
		{
			name: "Update priority",
			s:    client.GameServersSettings,
			args: args{
				serviceID: 7654321,
				opts: GSSettingsUpdateOptions{
					Category: "general",
					Key:      "priority",
					Value:    strings.Join([]string{"Gamer Tag 1", "Gamer Tag 2"}, "%0D%0A"),
				},
			},
			wantErr: false,
		},
		{
			name: "Update failure",
			s:    client.GameServersSettings,
			args: args{
				serviceID: 999,
				opts: GSSettingsUpdateOptions{
					Category: "general",
					Key:      "priority",
					Value:    "aa",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing category",
			s:    client.GameServersSettings,
			args: args{
				serviceID: 7654321,
				opts: GSSettingsUpdateOptions{
					Category: "",
					Key:      "priority",
					Value:    "aa",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing key",
			s:    client.GameServersSettings,
			args: args{
				serviceID: 7654321,
				opts: GSSettingsUpdateOptions{
					Category: "general",
					Key:      "",
					Value:    "aa",
				},
			},
			wantErr: true,
		},
		{
			name: "Missing value",
			s:    client.GameServersSettings,
			args: args{
				serviceID: 7654321,
				opts: GSSettingsUpdateOptions{
					Category: "general",
					Key:      "priority",
					Value:    "",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Update(tt.args.serviceID, tt.args.opts); (err != nil) != tt.wantErr {
				t.Errorf("GSSettingsService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

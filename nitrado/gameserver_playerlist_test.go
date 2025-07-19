package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// JSON minified using https://codebeautify.org/jsonminifier

// TestPlayerListService_List tests the PlayerListService List() method.
func TestPlayerListService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/games/players", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"players":[{"name":null,"id":"abcdefg12345678","id_type":"internal","online":"true","actions":["kick"]},{"name":"Player25","id":"hijklmnop0987654321","id_type":"internal","online":"true","actions":["kick","promotion_level_0","promotion_level_2","promotion_level_3"]},{"name":"Player26","id":"ijklmnop0987654322","id_type":"internal","online":"true","actions":["kick","promotion_level_0","promotion_level_2","promotion_level_3"]},{"name":null,"id":"bcdefg123456790","id_type":"internal","online":"true","actions":["kick"]}]}}`)
	})

	type args struct {
		svc Service
	}
	tests := []struct {
		name    string
		s       *PlayerListService
		args    args
		want    []Player
		wantErr bool
	}{
		{
			name: "PlayerList",
			s:    client.PlayerListService,
			args: args{
				svc: Service{ID: 7654321, Username: "ni1_1"},
			},
			want: []Player{
				{
					Name:    "",
					ID:      "abcdefg12345678",
					IDType:  "internal",
					Online:  "true",
					Actions: []string{"kick"},
				},
				{
					Name:    "",
					ID:      "bcdefg123456790",
					IDType:  "internal",
					Online:  "true",
					Actions: []string{"kick"},
				},
				{
					Name:    "Player25",
					ID:      "hijklmnop0987654321",
					IDType:  "internal",
					Online:  "true",
					Actions: []string{"kick", "promotion_level_0", "promotion_level_2", "promotion_level_3"},
				},
				{
					Name:    "Player26",
					ID:      "ijklmnop0987654322",
					IDType:  "internal",
					Online:  "true",
					Actions: []string{"kick", "promotion_level_0", "promotion_level_2", "promotion_level_3"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.List(tt.args.svc)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlayerListService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlayerListService.List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

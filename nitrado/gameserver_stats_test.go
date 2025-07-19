package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// JSON minified using https://codebeautify.org/jsonminifier

// TestGameserverStats_Get tests the GameServerStatsService List() method.
func TestGameserverStats_Get(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/stats", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"stats":{"cpuUsage":[[0,1622359920],[0,1622359980]],"currentPlayers":[[1,1622359920],[0,1622359980]],"maxPlayers":[[32,1622359920],[32,1622359980]],"memoryUsage":[[4732,1622359920],[4734,1622359980]]}}}`)
	})

	type args struct {
		serviceID int
	}
	tests := []struct {
		name    string
		s       *GameServerStatsService
		args    args
		want    GSStats
		wantErr bool
	}{
		{
			name: "DayZ Server Stats",
			s:    client.GameServerStats,
			args: args{
				serviceID: 7654321,
			},
			want: GSStats{
				CPUUsage: [][]float32{
					{0.0, 1622359920},
					{0.0, 1622359980},
				},
				CurrentPlayers: [][]float32{
					{1.0, 1622359920},
					{0.0, 1622359980},
				},
				MaxPlayers: [][]float32{
					{32.0, 1622359920},
					{32.0, 1622359980},
				},
				MemoryUsage: [][]float32{
					{4732.0, 1622359920},
					{4734.0, 1622359980},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.Get(tt.args.serviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GameServerStatsService.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("GameServerStatsService.Get() got = %v, want %v", got, &tt.want)
			}
		})
	}
}

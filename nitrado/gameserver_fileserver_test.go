package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestFileServerService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"success","data":{"entries":[{"owner":"ni1_1","chmod":"100664","size":499593,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64.ADM","accessed_at":1608633696,"group":"ni1_1","type":"file","created_at":1608633683,"modified_at":1608633679,"name":"DayZServer_X1_x64.ADM"},{"owner":"ni1_1","chmod":"100664","size":189313,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64_2020_12_21_165346976.ADM","accessed_at":1608566680,"group":"ni1_1","type":"file","created_at":1608566578,"modified_at":1608566026,"name":"DayZServer_X1_x64_2020_12_21_165346976.ADM"}]}}`)
	})

	type args struct {
		svc  Service
		opts FileServerListOptions
	}
	tests := []struct {
		name    string
		s       *FileServerService
		args    args
		want    *[]File
		wantErr bool
	}{
		{
			name: "DayZ Logs",
			s:    client.FileServerService,
			args: args{
				svc: Service{ID: 7654321, Username: "ni1_1"},
				opts: FileServerListOptions{
					Dir: "/games/ni1_1/noftp/dayzxb/config",
				},
			},
			want: &[]File{
				{
					Owner:      "ni1_1",
					Chmod:      "100664",
					Size:       189313,
					Path:       "/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64_2020_12_21_165346976.ADM",
					AccessedAt: 1608566680,
					Group:      "ni1_1",
					Type:       "file",
					CreatedAt:  1608566578,
					ModifiedAt: 1608566026,
					Name:       "DayZServer_X1_x64_2020_12_21_165346976.ADM",
				},
				{
					Owner:      "ni1_1",
					Chmod:      "100664",
					Size:       499593,
					Path:       "/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64.ADM",
					AccessedAt: 1608633696,
					Group:      "ni1_1",
					Type:       "file",
					CreatedAt:  1608633683,
					ModifiedAt: 1608633679,
					Name:       "DayZServer_X1_x64.ADM",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.List(tt.args.svc, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileServerService.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileServerService.List() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileServerService_Download(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"status":"success","data":{"token":{"url":"http://dev001.nitrado.net:8080/download/?token=00000000-0000-0000-0000-000000000000","token":"00000000-0000-0000-0000-000000000000"}}}`)
	})

	type args struct {
		svc  Service
		opts FileServerDownloadOptions
	}
	tests := []struct {
		name    string
		s       *FileServerService
		args    args
		want    string
		want1   *http.Response
		wantErr bool
	}{
		{
			name: "Download a file",
			s:    client.FileServerService,
			args: args{
				svc: Service{ID: 7654321, Username: "ni1_1"},
				opts: FileServerDownloadOptions{
					File: "abcd",
				},
			},
			want: "http://dev001.nitrado.net:8080/download/?token=00000000-0000-0000-0000-000000000000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.Download(tt.args.svc, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileServerService.Download() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, &tt.want) {
				t.Errorf("FileServerService.Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}

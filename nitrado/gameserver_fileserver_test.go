package nitrado

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// JSON minified using https://codebeautify.org/jsonminifier

// TestFileServerService_List tests the FileServerService List() method.
func TestFileServerService_List(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"entries":[{"owner":"ni1_1","chmod":"100664","size":499593,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64.ADM","accessed_at":1608633696,"group":"ni1_1","type":"file","created_at":1608633683,"modified_at":1608633679,"name":"DayZServer_X1_x64.ADM"},{"owner":"ni1_1","chmod":"100664","size":271029,"path":"/games/ni1_1/noftp/dayzxb/config/script_2021-07-31_14-23-42.log","accessed_at":1627745850,"group":"ni1_1","type":"file","created_at":1627745838,"modified_at":1627745795,"name":"script_2021-07-31_14-23-42.log"},{"owner":"ni1_1","chmod":"100664","size":189313,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64_2020_12_21_165346976.ADM","accessed_at":1608566680,"group":"ni1_1","type":"file","created_at":1608566578,"modified_at":1608566026,"name":"DayZServer_X1_x64_2020_12_21_165346976.ADM"}]}}`)
	})

	type args struct {
		svc  Service
		opts FileServerListOptions
	}
	tests := []struct {
		name    string
		s       *FileServerService
		args    args
		want    []File
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
			want: []File{
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
				{
					Owner:      "ni1_1",
					Chmod:      "100664",
					Size:       271029,
					Path:       "/games/ni1_1/noftp/dayzxb/config/script_2021-07-31_14-23-42.log",
					AccessedAt: 1627745850,
					Group:      "ni1_1",
					Type:       "file",
					CreatedAt:  1627745838,
					ModifiedAt: 1627745795,
					Name:       "script_2021-07-31_14-23-42.log",
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

// TestFileServerService_ListWithFilter tests the FileServerService List() method.
func TestFileServerService_ListWithFilter(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/list", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"entries":[{"owner":"ni1_1","chmod":"100664","size":499593,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64.ADM","accessed_at":1608633696,"group":"ni1_1","type":"file","created_at":1608633683,"modified_at":1608633679,"name":"DayZServer_X1_x64.ADM"},{"owner":"ni1_1","chmod":"100664","size":189313,"path":"/games/ni1_1/noftp/dayzxb/config/DayZServer_X1_x64_2020_12_21_165346976.ADM","accessed_at":1608566680,"group":"ni1_1","type":"file","created_at":1608566578,"modified_at":1608566026,"name":"DayZServer_X1_x64_2020_12_21_165346976.ADM"}]}}`)
	})

	type args struct {
		svc  Service
		opts FileServerListOptions
	}
	tests := []struct {
		name    string
		s       *FileServerService
		args    args
		want    []File
		wantErr bool
	}{
		{
			name: "DayZ Logs",
			s:    client.FileServerService,
			args: args{
				svc: Service{ID: 7654321, Username: "ni1_1"},
				opts: FileServerListOptions{
					Dir:    "/games/ni1_1/noftp/dayzxb/config",
					Search: "*.ADM",
				},
			},
			want: []File{
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

// TestFileServerService_Download tests the FileServerService Download() method.
func TestFileServerService_Download(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/download", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"token":{"url":"http://dev001.nitrado.net:8080/download/?token=00000000-0000-0000-0000-000000000000","token":"00000000-0000-0000-0000-000000000000"}}}`)
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
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileServerService.Download() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestFileServerService_Upload tests the FileServerService Download() method.
func TestFileServerService_Upload(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/services/7654321/gameservers/file_server/upload", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, _ = fmt.Fprint(w, `{"status":"success","data":{"token":{"url":"http://dev001.nitrado.net:8080/upload/","token":"00000000-0000-0000-0000-000000000000"}}}`)
	})

	type args struct {
		svc  Service
		opts FileServerUploadOptions
	}
	tests := []struct {
		name    string
		s       *FileServerService
		args    args
		want    FileDownloadResp
		wantErr bool
	}{
		{
			name: "Upload a file",
			s:    client.FileServerService,
			args: args{
				svc: Service{ID: 7654321, Username: "ni1_1"},
				opts: FileServerUploadOptions{
					File: "abcd.txt",
					Path: "/test",
				},
			},
			want: FileDownloadResp{
				Status: "success",
				Data: struct {
					Token struct {
						URL   string "json:\"url,omitempty\""
						Token string "json:\"token,omitempty\""
					} "json:\"token,omitempty\""
				}{
					Token: struct {
						URL   string "json:\"url,omitempty\""
						Token string "json:\"token,omitempty\""
					}{
						URL:   "http://dev001.nitrado.net:8080/upload/",
						Token: "00000000-0000-0000-0000-000000000000",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _, err := tt.s.Upload(tt.args.svc, tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileServerService.Upload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileServerService.Upload() got = %v, want %v", got, tt.want)
			}
		})
	}
}

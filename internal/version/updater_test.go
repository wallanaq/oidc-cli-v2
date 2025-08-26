package version

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestCheckForUpdate(t *testing.T) {
	setVersion := func(t *testing.T, v string) {
		originalVersion := version
		version = v
		t.Cleanup(func() {
			version = originalVersion
		})
	}

	setGhApiURL := func(t *testing.T, url string) {
		originalURL := ghApiURL
		ghApiURL = url
		t.Cleanup(func() {
			ghApiURL = originalURL
		})
	}

	testCases := []struct {
		name           string
		currentVersion string
		mockResponse   string
		mockStatus     int
		want           *UpdateInfo
		wantErr        bool
	}{
		{
			name:           "update available",
			currentVersion: "v1.0.0",
			mockResponse:   `{"tag_name": "v1.1.0"}`,
			mockStatus:     http.StatusOK,
			want: &UpdateInfo{
				UpdateAvailable: true,
				CurrentVersion:  "v1.0.0",
				LatestVersion:   "v1.1.0",
			},
			wantErr: false,
		},
		{
			name:           "no update available (same version)",
			currentVersion: "v1.1.0",
			mockResponse:   `{"tag_name": "v1.1.0"}`,
			mockStatus:     http.StatusOK,
			want: &UpdateInfo{
				UpdateAvailable: false,
				CurrentVersion:  "v1.1.0",
				LatestVersion:   "v1.1.0",
			},
			wantErr: false,
		},
		{
			name:           "no update available (local is newer)",
			currentVersion: "v1.2.0",
			mockResponse:   `{"tag_name": "v1.1.0"}`,
			mockStatus:     http.StatusOK,
			want: &UpdateInfo{
				UpdateAvailable: false,
				CurrentVersion:  "v1.2.0",
				LatestVersion:   "v1.1.0",
			},
			wantErr: false,
		},
		{
			name:           "dev version should not check for updates",
			currentVersion: "dev",
			want: &UpdateInfo{
				UpdateAvailable: false,
				CurrentVersion:  "dev",
			},
			wantErr: false,
		},
		{
			name:           "invalid semver should not check for updates",
			currentVersion: "not-a-semver",
			want: &UpdateInfo{
				UpdateAvailable: false,
				CurrentVersion:  "not-a-semver",
			},
			wantErr: false,
		},
		{
			name:           "github api returns non-200 status",
			currentVersion: "v1.0.0",
			mockStatus:     http.StatusInternalServerError,
			wantErr:        true,
		},
		{
			name:           "github api returns invalid json",
			currentVersion: "v1.0.0",
			mockResponse:   `{`,
			mockStatus:     http.StatusOK,
			wantErr:        true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.mockStatus)
				fmt.Fprintln(w, tc.mockResponse)
			}))
			defer server.Close()

			setVersion(t, tc.currentVersion)

			if tc.mockStatus != 0 {
				setGhApiURL(t, server.URL)
			}

			got, err := CheckForUpdate(context.Background())

			if (err != nil) != tc.wantErr {
				t.Fatalf("CheckForUpdate() error = %v, wantErr %v", err, tc.wantErr)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("CheckForUpdate() got = %v, want %v", got, tc.want)
			}
		})
	}
}

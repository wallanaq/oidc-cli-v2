package version

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"golang.org/x/mod/semver"
)

var ghApiURL = "https://api.github.com/repos/wallanaq/oidc-cli-v2/releases/latest"

type Release struct {
	TagName string `json:"tag_name"`
}

type UpdateInfo struct {
	UpdateAvailable bool
	CurrentVersion  string
	LatestVersion   string
}

func CheckForUpdate(ctx context.Context) (*UpdateInfo, error) {
	if version == "dev" || !semver.IsValid(version) {
		return &UpdateInfo{UpdateAvailable: false, CurrentVersion: version}, nil
	}

	slog.Debug("Checking for updates", slog.String("githubReleaseURL", ghApiURL))

	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ghApiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "oidc-cli")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &UpdateInfo{
		UpdateAvailable: semver.Compare(version, release.TagName) < 0,
		CurrentVersion:  version,
		LatestVersion:   release.TagName,
	}, nil
}

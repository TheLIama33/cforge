package updater

import (
	"fmt"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

type UpdateResult struct {
	Updated       bool
	LatestVersion string
	ReleaseNotes  string
}

func CheckAndUpdate(currentVersionStr string, repoSlug string) (*UpdateResult, error) {
	if currentVersionStr == "dev" {
		return nil, fmt.Errorf("cannot update a development build")
	}

	v, err := semver.Parse(currentVersionStr)
	if err != nil {
		return nil, fmt.Errorf("invalid current version syntax '%s': %w", currentVersionStr, err)
	}

	fmt.Println("Checking for updates...")
	latest, found, err := selfupdate.DetectLatest(repoSlug)
	if err != nil {
		return nil, fmt.Errorf("error checking GitHub for updates: %w", err)
	}

	if !found {
		return nil, fmt.Errorf("no release found for %s", repoSlug)
	}

	if latest.Version.LTE(v) {
		return &UpdateResult{
			Updated:       false,
			LatestVersion: latest.Version.String(),
		}, nil
	}

	fmt.Printf("Found new version %s. Downloading...\n", latest.Version)

	exe, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("could not locate executable path: %w", err)
	}

	if err := selfupdate.UpdateTo(latest.AssetURL, exe); err != nil {
		return nil, fmt.Errorf("binary update failed: %w", err)
	}

	return &UpdateResult{
		Updated:       true,
		LatestVersion: latest.Version.String(),
		ReleaseNotes:  latest.ReleaseNotes,
	}, nil
}

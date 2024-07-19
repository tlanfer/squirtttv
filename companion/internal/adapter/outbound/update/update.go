package update

import (
	"context"
	"fmt"
	"github.com/creativeprojects/go-selfupdate"
	"log"
)

var Version string

const repoSlug = "tlanfer/squirtttv"

func IsLatest() (string, bool, error) {
	latest, found, err := selfupdate.DetectLatest(context.Background(), selfupdate.ParseSlug(repoSlug))

	if err != nil {
		return "", false, fmt.Errorf("failed to detect latest version: %v", err)
	}

	if !found {
		return "", false, fmt.Errorf("failed to find latest version")
	}

	return latest.Version(), latest.LessOrEqual(Version), nil
}

func Run() error {
	latest, found, err := selfupdate.DetectLatest(context.Background(), selfupdate.ParseSlug(repoSlug))

	if err != nil {
		return fmt.Errorf("failed to detect latest version: %v", err)
	}

	if !found {
		return fmt.Errorf("failed to find latest version")
	}

	if latest.LessOrEqual(Version) {
		log.Printf("current version %s is up-to-date", Version)
		return nil
	}

	exe, err := selfupdate.ExecutablePath()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %v", err)
	}

	err = selfupdate.UpdateTo(context.Background(), latest.AssetURL, latest.AssetName, exe)
	if err != nil {
		return fmt.Errorf("failed to update: %v", err)
	}
	log.Printf("successfully updated to version %v", latest.Version())

	return nil
}

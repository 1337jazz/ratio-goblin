package updater

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	goupdate "github.com/inconshreveable/go-update"

	"github.com/1337jazz/ratio-goblin/internal/version"
)

const releaseURL = "https://api.github.com/repos/1337jazz/ratio-goblin/releases/latest"

var UpdateCancelled = errors.New("update cancelled by user")
var AlreadyUpToDate = errors.New("already up to date")

type release struct {
	TagName string `json:"tag_name"`
	Body    string `json:"body"`
}

// fetchRelease retrieves the latest release information from GitHub, respecting a 3-second timeout
func fetchRelease() (*release, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, releaseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned %d", resp.StatusCode)
	}

	var r release
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}

	return &r, nil
}

// HasUpdate returns true if a newer release exists on GitHub, fails silently (returns false) on any error or timeout
func HasUpdate() bool {
	r, err := fetchRelease()
	if err != nil {
		return false
	}
	latest := strings.TrimPrefix(r.TagName, "v")
	current := strings.TrimPrefix(version.Version, "v")
	return latest != "" && latest != current
}

// Update fetches release info, prompts the user with version and changelog,
// and applies the update if confirmed
func Update() error {
	r, err := fetchRelease()
	if err != nil {
		return fmt.Errorf("failed to fetch latest release: %w", err)
	}

	// Compare the current and latest versions, ignoring any leading "v"
	latest := strings.TrimPrefix(r.TagName, "v")
	current := strings.TrimPrefix(version.Version, "v")
	if latest == current {
		return AlreadyUpToDate
	}

	// Ask the user to confirm the update, showing the changelog if available
	fmt.Printf("\nUpdate available: %s -> %s\n\n", current, latest)
	if r.Body != "" {
		fmt.Printf("%s\n", r.Body)
	}
	fmt.Print("Update now? [y/N] ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		return UpdateCancelled
	}

	tarURL := fmt.Sprintf(
		"https://github.com/1337jazz/ratio-goblin/releases/download/%s/ratiogoblin_Linux_x86_64.tar.gz",
		r.TagName,
	)

	// Get the tarball from GitHub and apply the update
	resp, err := http.Get(tarURL) //nolint:noctx
	if err != nil {
		return fmt.Errorf("failed to download release: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download returned %d", resp.StatusCode)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to decompress release: %w", err)
	}
	defer gz.Close()

	// Iterate through the tarball to find the binary and apply the update (safe to assume it's named "ratiogoblin")
	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read archive: %w", err)
		}
		if hdr.Name == "ratiogoblin" {
			if err := goupdate.Apply(tr, goupdate.Options{}); err != nil {
				if rerr := goupdate.RollbackError(err); rerr != nil {
					return fmt.Errorf("update failed and rollback failed: %w", rerr)
				}
				return fmt.Errorf("update failed (rolled back): %w", err)
			}
			return nil
		}
	}

	return fmt.Errorf("binary not found in release archive")
}

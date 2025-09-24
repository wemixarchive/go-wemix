package solcdownloader

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	solcBinDirName = ".solc-bin"                                                                   // Name of the local directory where solc binaries are stored
	solcRepoURL    = "https://binaries.soliditylang.org"                                           // Base URL of the official solc binaries repository
	osMapping      = map[string]string{"linux": "linux", "darwin": "macosx", "windows": "windows"} // Mapping from runtime.GOOS to solc binary OS names
)

// Build represents a single solc binary build from list.json
type Build struct {
	Path        string   `json:"path"`
	Version     string   `json:"version"`
	Build       string   `json:"build"`
	LongVersion string   `json:"longVersion"`
	Sha256      string   `json:"sha256"`
	Urls        []string `json:"urls"`
}

// ListJSON represents the list.json structure from solc releases
type ListJSON struct {
	Builds []Build `json:"builds"`
}

// GetSolcBin returns the path to the local solc binary for the specified version.
// If not present, it downloads the binary, verifies its checksum, and sets execution permission.
func GetSolcBin(version string) (string, error) {
	osName, ok := osMapping[runtime.GOOS]
	if !ok {
		return "", fmt.Errorf("unsupported OS: %s", runtime.GOOS)
	}
	arch := runtime.GOARCH
	// Handle ARM architecture
	if runtime.GOARCH == "arm64" {
		if osName == "macosx" {
			// Mac ARM → check Rosetta 2
			if err := exec.Command("pgrep", "oahd").Run(); err != nil {
				return "", fmt.Errorf("rosetta 2 is required to run amd64 solc binary on Apple Silicon")
			}
			arch = "amd64"
		} else {
			// ARM on non-Mac OS → unsupported
			return "", fmt.Errorf("unsupported architecture: ARM is not supported on %s", osName)
		}
	}

	// Home directory based path for storing solc binaries
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, solcBinDirName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	// Local binary file name
	binName := fmt.Sprintf("solc-%s", version)
	if osName == "windows" {
		binName += ".exe"
	}
	binPath := filepath.Join(dir, binName)

	// Return the path immediately if the binary already exists
	if _, err := os.Stat(binPath); err == nil {
		return binPath, nil
	}

	fmt.Printf("%s not found, downloading...\n", binName)

	// Windows versions below 0.7.2 are unsupported
	if osName == "windows" && versionCompare(version, "0.7.2") < 0 {
		return "", fmt.Errorf("unsupported version on Windows: %s. Please use >=0.7.2", version)
	}

	// Fetch list.json to find the build info
	listURL := fmt.Sprintf("%s/%s-%s/list.json", solcRepoURL, osName, arch)
	build, err := getBuild(listURL, version)
	if err != nil {
		return "", err
	}

	// Construct download URL using the build path
	downloadUrl := fmt.Sprintf("%s/%s-%s/%s", solcRepoURL, osName, arch, build.Path)
	if err := downloadSolc(binPath, downloadUrl); err != nil {
		return "", err
	}

	// Verify SHA256 checksum
	if ok, err := verifyChecksum(binPath, build.Sha256); err != nil {
		return "", err
	} else if !ok {
		os.Remove(binPath)
		return "", fmt.Errorf("downloaded file checksum mismatch")
	}

	// Set executable permission on macOS/Linux
	if osName != "windows" {
		if err := os.Chmod(binPath, 0755); err != nil {
			return "", err
		}
	}

	fmt.Println("Download complete:", binPath)
	return binPath, nil
}

// getBuild fetches the list.json and returns the Build corresponding to the requested version
func getBuild(listURL, version string) (*Build, error) {
	resp, err := http.Get(listURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch list.json: %s", resp.Status)
	}

	var data ListJSON
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	// Find the requested version in the builds list
	for _, build := range data.Builds {
		if build.Version == version {
			return &build, nil
		}
	}
	return nil, fmt.Errorf("version %s not found in list.json", version)
}

// downloadSolc downloads a file from the given URL and writes it to the specified path.
// It uses a temporary file and moves it to the final location after a successful download.
func downloadSolc(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	tmpFile := filepath + ".tmp"
	out, err := os.Create(tmpFile)
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		os.Remove(tmpFile)
		return err
	}

	// Rename the temporary file to the final path
	return os.Rename(tmpFile, filepath)
}

// verifyChecksum calculates SHA256 of the file and compares it with the expected hash
func verifyChecksum(filepath string, sha256Hex string) (bool, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	calculated := "0x" + hex.EncodeToString(h.Sum(nil))
	return calculated == sha256Hex, nil
}

// versionCompare compares two semantic versions.
// Returns -1 if v1<v2, 0 if v1==v2, 1 if v1>v2
func versionCompare(v1, v2 string) int {
	v1 = strings.TrimPrefix(v1, "v")
	v2 = strings.TrimPrefix(v2, "v")

	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var n1, n2 int
		if i < len(parts1) {
			fmt.Sscanf(parts1[i], "%d", &n1)
		}
		if i < len(parts2) {
			fmt.Sscanf(parts2[i], "%d", &n2)
		}
		if n1 < n2 {
			return -1
		} else if n1 > n2 {
			return 1
		}
	}
	return 0
}

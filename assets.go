package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0o755)
	}
	return nil
}

// getExtension extracts and cleans the file extension from a media type
func getExtension(mediaType string) (string, error) {
	parsedType, _, err := mime.ParseMediaType(mediaType)
	if err != nil {
		return "", err
	}

	extensions, err := mime.ExtensionsByType(parsedType)
	if err != nil || len(extensions) == 0 {
		return "", fmt.Errorf("unsupported media type")
	}

	return strings.TrimPrefix(extensions[0], "."), nil
}

// getAssetPath retains the original behavior for videos named by UUID
// func getAssetPath(videoID uuid.UUID, mediaType string) (string, error) {
// 	ext, err := getExtension(mediaType)
// 	if err != nil {
// 		return "", err
// 	}
// 	return fmt.Sprintf("%s.%s", videoID, ext), nil
// }

// getRandomAssetPath generates a random base64 filename for thumbnails
func getRandomAssetPath(mediaType string) (string, error) {
	base := make([]byte, 32)
	if _, err := rand.Read(base); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	id := base64.RawURLEncoding.EncodeToString(base)

	ext, err := getExtension(mediaType)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s", id, ext), nil
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

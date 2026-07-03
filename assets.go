package main

import (
	"fmt"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (cfg apiConfig) ensureAssetsDir() error {
	if _, err := os.Stat(cfg.assetsRoot); os.IsNotExist(err) {
		return os.Mkdir(cfg.assetsRoot, 0o755)
	}
	return nil
}

func getAssetPath(videoID uuid.UUID, mediaType string) (string, error) {
	// Parse the media type to strip out parameters like charset
	parsedType, _, err := mime.ParseMediaType(mediaType)
	if err != nil {
		return "", err
	}

	// Look up valid extensions for the media type
	extensions, err := mime.ExtensionsByType(parsedType)
	if err != nil || len(extensions) == 0 {
		return "", fmt.Errorf("unsupported media type")
	}

	// Clean the leading dot from the extension
	ext := strings.TrimPrefix(extensions[0], ".")
	return fmt.Sprintf("%s.%s", videoID, ext), nil
}

func (cfg apiConfig) getAssetDiskPath(assetPath string) string {
	return filepath.Join(cfg.assetsRoot, assetPath)
}

func (cfg apiConfig) getAssetURL(assetPath string) string {
	return fmt.Sprintf("http://localhost:%s/assets/%s", cfg.port, assetPath)
}

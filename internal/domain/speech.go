package domain

import (
	"fmt"
	"path/filepath"
	"strings"
)

type OutputFormat int

const (
	FormatPlay OutputFormat = iota
	FormatAIFF
	FormatWAV
	FormatM4A
	FormatMP3
)

func DetectFormat(path string) (OutputFormat, error) {
	if path == "" {
		return FormatPlay, nil
	}
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".aiff", ".aif":
		return FormatAIFF, nil
	case ".wav":
		return FormatWAV, nil
	case ".m4a", ".aac":
		return FormatM4A, nil
	case ".mp3":
		return FormatMP3, nil
	default:
		return 0, fmt.Errorf("非対応の拡張子: %q (対応: .aiff/.wav/.m4a/.aac/.mp3)", ext)
	}
}

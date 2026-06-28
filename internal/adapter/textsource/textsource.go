package textsource

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

type Args struct{ Values []string }

func (s *Args) ReadText(_ context.Context) (string, error) {
	if len(s.Values) == 0 {
		return "", nil
	}
	return strings.Join(s.Values, " "), nil
}

type File struct{ Path string }

func (s *File) ReadText(_ context.Context) (string, error) {
	if s.Path == "-" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", fmt.Errorf("標準入力の読み込みに失敗: %w", err)
		}
		return string(b), nil
	}
	b, err := os.ReadFile(s.Path)
	if err != nil {
		return "", fmt.Errorf("ファイルの読み込みに失敗: %w", err)
	}
	return string(b), nil
}

type Stdin struct{}

func (s *Stdin) ReadText(_ context.Context) (string, error) {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("標準入力の読み込みに失敗: %w", err)
	}
	return string(b), nil
}

type Clipboard struct{}

func (s *Clipboard) ReadText(ctx context.Context) (string, error) {
	if _, err := exec.LookPath("pbpaste"); err != nil {
		return "", fmt.Errorf("クリップボード読み取りに pbpaste が必要です")
	}
	out, err := exec.CommandContext(ctx, "pbpaste").Output()
	if err != nil {
		return "", fmt.Errorf("pbpaste の実行に失敗: %w", err)
	}
	return string(out), nil
}

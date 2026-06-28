package ffmpeg

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

type Encoder struct{}

func New() *Encoder { return &Encoder{} }

func (e *Encoder) Encode(ctx context.Context, srcAIFF, dst string) error {
	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("MP3 出力には ffmpeg が必要です: brew install ffmpeg")
	}
	cmd := exec.CommandContext(ctx, "ffmpeg", "-y", "-loglevel", "error", "-i", srcAIFF,
		"-codec:a", "libmp3lame", "-qscale:a", "2", dst)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

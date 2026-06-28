package say

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/junara/jmactts/internal/domain"
)

type Adapter struct{}

func New() *Adapter { return &Adapter{} }

func (a *Adapter) Play(ctx context.Context, text, voice string, rate int) error {
	return a.run(ctx, text, voice, rate, nil)
}

func (a *Adapter) Save(ctx context.Context, text, voice string, rate int, path string, format domain.OutputFormat) error {
	out := &output{path: path}
	switch format {
	case domain.FormatAIFF:
		// say default
	case domain.FormatWAV:
		out.fileFormat = "WAVE"
		out.dataFormat = "LEI16@22050"
	case domain.FormatM4A:
		out.fileFormat = "m4af"
		out.dataFormat = "aac"
	default:
		return fmt.Errorf("say adapter: 非対応のフォーマット (%d)", format)
	}
	return a.run(ctx, text, voice, rate, out)
}

type output struct {
	path       string
	fileFormat string
	dataFormat string
}

func (a *Adapter) run(ctx context.Context, text, voice string, rate int, out *output) error {
	args := []string{}
	if voice != "" {
		args = append(args, "-v", voice)
	}
	if rate > 0 {
		args = append(args, "-r", fmt.Sprintf("%d", rate))
	}
	if out != nil {
		args = append(args, "-o", out.path)
		if out.fileFormat != "" {
			args = append(args, "--file-format="+out.fileFormat)
		}
		if out.dataFormat != "" {
			args = append(args, "--data-format="+out.dataFormat)
		}
	}
	cmd := exec.CommandContext(ctx, "say", args...)
	cmd.Stdin = strings.NewReader(text)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var voiceLineRE = regexp.MustCompile(`^(.+?)\s+([a-z]{2})_([A-Z]{2})\s+#\s*(.*)$`)

func (a *Adapter) List(ctx context.Context) (domain.VoiceList, error) {
	out, err := exec.CommandContext(ctx, "say", "-v", "?").Output()
	if err != nil {
		return nil, fmt.Errorf("say -v ? の実行に失敗: %w", err)
	}
	var voices domain.VoiceList
	for _, line := range strings.Split(string(out), "\n") {
		m := voiceLineRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		voices = append(voices, domain.Voice{
			Name:        strings.TrimSpace(m[1]),
			Locale:      domain.Locale{Lang: m[2], Country: m[3]},
			Description: strings.TrimSpace(m[4]),
		})
	}
	return voices, nil
}

package usecase

import (
	"context"

	"github.com/junara/jmactts/internal/domain"
)

type Synthesizer interface {
	Play(ctx context.Context, text, voice string, rate int) error
	Save(ctx context.Context, text, voice string, rate int, path string, format domain.OutputFormat) error
}

type MP3Encoder interface {
	Encode(ctx context.Context, srcAIFF, dst string) error
}

type VoiceCatalog interface {
	List(ctx context.Context) (domain.VoiceList, error)
}

type TextSource interface {
	ReadText(ctx context.Context) (string, error)
}

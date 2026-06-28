package usecase

import (
	"context"
	"os"

	"github.com/junara/jmactts/internal/domain"
)

const ChunkSize = 1500

type SpeakUseCase struct {
	Synth   Synthesizer
	Encoder MP3Encoder
}

type SpeakInput struct {
	Text       string
	Voice      string
	Rate       int
	OutputPath string
	Format     domain.OutputFormat
}

func (uc *SpeakUseCase) Execute(ctx context.Context, in SpeakInput) error {
	if in.Format == domain.FormatPlay {
		return uc.play(ctx, in)
	}
	return uc.save(ctx, in)
}

func (uc *SpeakUseCase) play(ctx context.Context, in SpeakInput) error {
	for _, chunk := range domain.ChunkText(in.Text, ChunkSize) {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := uc.Synth.Play(ctx, chunk, in.Voice, in.Rate); err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			return err
		}
	}
	return nil
}

func (uc *SpeakUseCase) save(ctx context.Context, in SpeakInput) error {
	if in.Format != domain.FormatMP3 {
		return uc.Synth.Save(ctx, in.Text, in.Voice, in.Rate, in.OutputPath, in.Format)
	}
	tmp, err := os.CreateTemp("", "jmactts-*.aiff")
	if err != nil {
		return err
	}
	tmpPath := tmp.Name()
	tmp.Close()
	defer os.Remove(tmpPath)

	if err := uc.Synth.Save(ctx, in.Text, in.Voice, in.Rate, tmpPath, domain.FormatAIFF); err != nil {
		return err
	}
	return uc.Encoder.Encode(ctx, tmpPath, in.OutputPath)
}

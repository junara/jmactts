package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/junara/jmactts/internal/adapter/cli"
	"github.com/junara/jmactts/internal/adapter/ffmpeg"
	"github.com/junara/jmactts/internal/adapter/say"
	"github.com/junara/jmactts/internal/usecase"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	sayAdapter := say.New()
	mp3Encoder := ffmpeg.New()

	deps := cli.Deps{
		Speak: &usecase.SpeakUseCase{
			Synth:   sayAdapter,
			Encoder: mp3Encoder,
		},
		ListVoices: &usecase.ListVoices{Catalog: sayAdapter},
		PickVoice:  &usecase.PickVoice{Catalog: sayAdapter},
	}

	os.Exit(cli.Run(ctx, os.Args[1:], deps, os.Stdout, os.Stderr))
}

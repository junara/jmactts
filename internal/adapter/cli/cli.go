package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/junara/jmactts/internal/adapter/textsource"
	"github.com/junara/jmactts/internal/domain"
	"github.com/junara/jmactts/internal/usecase"
)

// Version is the build version, populated via -ldflags by goreleaser.
var Version = "dev"

type Deps struct {
	Speak      *usecase.SpeakUseCase
	ListVoices *usecase.ListVoices
	PickVoice  *usecase.PickVoice
}

func Run(ctx context.Context, args []string, deps Deps, stdout, stderr io.Writer) int {
	opts, code := parseFlags(args, stderr)
	if code >= 0 {
		return code
	}

	if opts.help {
		PrintUsage(stdout)
		return 0
	}

	if opts.showVersion {
		fmt.Fprintf(stdout, "jmactts %s\n", Version)
		return 0
	}

	if opts.listVoices {
		return runListVoices(ctx, deps, opts.lang, stdout, stderr)
	}

	if opts.voice == "" && opts.lang != "" {
		v, err := deps.PickVoice.Execute(ctx, opts.lang)
		if err != nil {
			fmt.Fprintf(stderr, "jmactts: %v\n", err)
			return 1
		}
		opts.voice = v
	}

	src := resolveSource(opts.useClipboard, opts.inputFile, opts.posArgs)
	text, err := src.ReadText(ctx)
	if err != nil {
		fmt.Fprintf(stderr, "jmactts: %v\n", err)
		return 1
	}
	if strings.TrimSpace(text) == "" {
		PrintUsage(stderr)
		return 2
	}

	format, err := domain.DetectFormat(opts.output)
	if err != nil {
		fmt.Fprintf(stderr, "jmactts: %v\n", err)
		return 1
	}

	err = deps.Speak.Execute(ctx, usecase.SpeakInput{
		Text:       text,
		Voice:      opts.voice,
		Rate:       opts.rate,
		OutputPath: opts.output,
		Format:     format,
	})
	if err != nil {
		if errors.Is(err, context.Canceled) || ctx.Err() != nil {
			return 130
		}
		fmt.Fprintf(stderr, "jmactts: %v\n", err)
		return 1
	}
	return 0
}

type options struct {
	voice        string
	lang         string
	rate         int
	inputFile    string
	output       string
	useClipboard bool
	listVoices   bool
	help         bool
	showVersion  bool
	posArgs      []string
}

func parseFlags(args []string, stderr io.Writer) (options, int) {
	var opts options
	fs := flag.NewFlagSet("jmactts", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() { PrintUsage(stderr) }

	fs.StringVar(&opts.voice, "v", "", "")
	fs.StringVar(&opts.voice, "voice", "", "")
	fs.StringVar(&opts.lang, "L", "", "")
	fs.StringVar(&opts.lang, "lang", "", "")
	fs.IntVar(&opts.rate, "r", 0, "")
	fs.IntVar(&opts.rate, "rate", 0, "")
	fs.StringVar(&opts.inputFile, "f", "", "")
	fs.StringVar(&opts.inputFile, "file", "", "")
	fs.StringVar(&opts.output, "o", "", "")
	fs.StringVar(&opts.output, "output", "", "")
	fs.BoolVar(&opts.useClipboard, "c", false, "")
	fs.BoolVar(&opts.useClipboard, "clipboard", false, "")
	fs.BoolVar(&opts.listVoices, "l", false, "")
	fs.BoolVar(&opts.listVoices, "list-voices", false, "")
	fs.BoolVar(&opts.help, "h", false, "")
	fs.BoolVar(&opts.help, "help", false, "")
	fs.BoolVar(&opts.showVersion, "V", false, "")
	fs.BoolVar(&opts.showVersion, "version", false, "")

	if err := fs.Parse(args); err != nil {
		return opts, 2
	}
	opts.posArgs = fs.Args()
	return opts, -1
}

func runListVoices(ctx context.Context, deps Deps, lang string, stdout, stderr io.Writer) int {
	voices, err := deps.ListVoices.Execute(ctx, lang)
	if err != nil {
		fmt.Fprintf(stderr, "jmactts: %v\n", err)
		return 1
	}
	for _, v := range voices {
		if lang == "" && v.Description != "" {
			fmt.Fprintf(stdout, "%-30s %s    # %s\n", v.Name, v.Locale.String(), v.Description)
		} else {
			fmt.Fprintf(stdout, "%-30s %s\n", v.Name, v.Locale.String())
		}
	}
	return 0
}

func resolveSource(useClipboard bool, file string, args []string) usecase.TextSource {
	if useClipboard {
		return &textsource.Clipboard{}
	}
	if file != "" {
		return &textsource.File{Path: file}
	}
	if len(args) > 0 {
		return &textsource.Args{Values: args}
	}
	if stat, err := os.Stdin.Stat(); err == nil && (stat.Mode()&os.ModeCharDevice) == 0 {
		return &textsource.Stdin{}
	}
	return &textsource.Args{}
}

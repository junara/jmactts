package usecase

import (
	"context"
	"fmt"

	"github.com/junara/jmactts/internal/domain"
)

type ListVoices struct {
	Catalog VoiceCatalog
}

func (uc *ListVoices) Execute(ctx context.Context, query string) (domain.VoiceList, error) {
	voices, err := uc.Catalog.List(ctx)
	if err != nil {
		return nil, err
	}
	if query == "" {
		return voices, nil
	}
	matches := voices.FilterByQuery(query)
	if len(matches) == 0 {
		return nil, fmt.Errorf("%q に一致するボイスが見つかりません", query)
	}
	return matches, nil
}

type PickVoice struct {
	Catalog VoiceCatalog
}

func (uc *PickVoice) Execute(ctx context.Context, query string) (string, error) {
	voices, err := uc.Catalog.List(ctx)
	if err != nil {
		return "", err
	}
	matches := voices.FilterByQuery(query)
	primary, ok := matches.Primary()
	if !ok {
		return "", fmt.Errorf("%q に一致するボイスが見つかりません (jmactts -l で一覧確認)", query)
	}
	return primary.Name, nil
}

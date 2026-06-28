package domain

import "strings"

func ChunkText(text string, maxRunes int) []string {
	runes := []rune(text)
	if len(runes) <= maxRunes {
		return []string{text}
	}
	terminators := map[rune]bool{
		'。': true, '．': true, '.': true,
		'!': true, '?': true, '！': true, '？': true,
		'\n': true,
	}
	var chunks []string
	start := 0
	for i := 0; i < len(runes); i++ {
		if i-start+1 >= maxRunes && terminators[runes[i]] {
			chunks = append(chunks, strings.TrimSpace(string(runes[start:i+1])))
			start = i + 1
		}
	}
	if start < len(runes) {
		rest := strings.TrimSpace(string(runes[start:]))
		if rest != "" {
			chunks = append(chunks, rest)
		}
	}
	return chunks
}

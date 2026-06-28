package domain

import (
	"strings"
	"testing"
)

func TestChunkText_NoSplitUnderThreshold(t *testing.T) {
	got := ChunkText("短い文章です。", 100)
	if len(got) != 1 {
		t.Fatalf("want 1 chunk, got %d", len(got))
	}
}

func TestChunkText_SplitsOnJapanesePeriod(t *testing.T) {
	sentence := strings.Repeat("あ", 50) + "。"
	text := strings.Repeat(sentence, 6) // 306 runes
	chunks := ChunkText(text, 100)
	if len(chunks) < 2 {
		t.Fatalf("want >=2 chunks, got %d", len(chunks))
	}
	if strings.Join(chunks, "") != strings.ReplaceAll(text, "\n", "") {
		// note: chunkText trims spaces; trailing newlines are stripped, but our text has none
		joined := strings.Join(chunks, "")
		if joined != text {
			t.Fatalf("chunks do not reconstruct input; joined len=%d orig len=%d", len([]rune(joined)), len([]rune(text)))
		}
	}
}

func TestChunkText_SplitsOnNewline(t *testing.T) {
	text := strings.Repeat("A", 200) + "\n" + strings.Repeat("B", 200)
	chunks := ChunkText(text, 100)
	if len(chunks) != 2 {
		t.Fatalf("want 2 chunks, got %d: %v", len(chunks), chunks)
	}
	if !strings.HasPrefix(chunks[0], "A") || !strings.HasPrefix(chunks[1], "B") {
		t.Fatalf("unexpected chunk content: %v", chunks)
	}
}

func TestChunkText_EnglishPunct(t *testing.T) {
	text := strings.Repeat("Hello world. ", 50) // ~650 chars, ".": every 13 chars
	chunks := ChunkText(text, 100)
	if len(chunks) < 2 {
		t.Fatalf("want >=2 chunks, got %d", len(chunks))
	}
}

---
title: Output Formats
description: Save audio as AIFF / WAV / M4A / AAC / MP3
---

Use `-o` / `--output` to save to a file. The format is **detected from the extension**.

| Extension | Format | Path |
|---|---|---|
| `.aiff` / `.aif` | AIFF (PCM, `say` default) | `say` directly |
| `.wav` | WAVE / LEI16 22050 Hz mono | `say` directly |
| `.m4a` / `.aac` | AAC in M4A | `say` directly |
| `.mp3` | MPEG Layer III (libmp3lame VBR Q2) | temp AIFF → `ffmpeg` |

## Examples

```bash
jmactts -L en -o hello.aiff "Hello"
jmactts -L en -o hello.wav  "Hello"
jmactts -L en -o hello.m4a  "Hello"
jmactts -L en -o hello.mp3  "Hello"        # requires ffmpeg
```

## How MP3 output works

`say` can't produce MP3 directly, so `jmactts` does the following internally:

1. Synthesize a temporary AIFF with `say`
2. Transcode to MP3 via `ffmpeg -codec:a libmp3lame -qscale:a 2`
3. Remove the temp file

MP3 output therefore **requires `ffmpeg`**:

```bash
brew install ffmpeg
```

Without it you'll see:

```
jmactts: MP3 出力には ffmpeg が必要です: brew install ffmpeg
```

## Notes on file output

- Long input is **not split** when writing to a file; you get a single audio file (see [Long Text & Interrupt](/jmactts/en/usage/long-text/))
- Existing files are overwritten silently

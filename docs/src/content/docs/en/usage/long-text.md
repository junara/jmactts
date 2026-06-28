---
title: Long Text & Interrupt
description: Automatic chunking for long input and instant Ctrl-C cancellation
---

## Automatic chunking

In playback mode (when `-o` is not given), **input over 1500 runes is split on sentence boundaries** and fed to `say` chunk by chunk. This keeps `say` stable on very long input and makes `Ctrl-C` responsive.

- Boundaries: `。` `．` `.` `!` `?` `！` `？` newline
- Each chunk extends past 1500 runes until the next boundary

### Example

```bash
cat long_article.txt | jmactts -L en
```

A 5000-rune article ends up as 3–5 chunks played back in order. A brief silence appears between chunks.

### File output isn't chunked

When writing to a file with `-o`, jmactts hands the whole text to `say` and produces a single audio file (chunking would create multiple files).

## Ctrl-C cancellation

Pressing `Ctrl-C` (SIGINT) during playback immediately kills the current chunk's `say` process and skips the rest.

- Exit code: **130** (= 128 + SIGINT)
- `SIGTERM` behaves the same way

### How it works

- A `context.Context` from `signal.NotifyContext` is passed to `exec.CommandContext`, so receiving a signal kills `say`
- The chunk loop checks `ctx.Err()` between chunks and exits when canceled

Long-running playback stops within ~1 second.

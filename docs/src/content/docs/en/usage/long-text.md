---
title: Long Text & Interrupt
description: Automatic chunking for long input and instant Ctrl-C cancellation
---

## Automatic chunking

When the input exceeds 1500 runes (live playback only), jmactts splits it at sentence boundaries and feeds chunks to `say` sequentially.

- Boundaries: `。` `．` `.` `!` `?` `！` `？` newline
- Each chunk includes everything up to the next boundary after reaching ~1500 runes

### Example

```bash
cat long_article.txt | jmactts -L en
```

A 5000-rune article ends up split into 3–4 chunks played back in order.

### File output behaviour

When writing to a file via `-o`, jmactts **does not chunk**. The entire text is handed to `say` and a single audio file is produced.

## Ctrl-C cancellation

Pressing `Ctrl-C` (`SIGINT`) during playback immediately kills the current chunk's `say` process and stops the loop.

- Exit code: **130** (= 128 + SIGINT)
- `SIGTERM` is handled the same way

### How it works

- A `context.Context` is created via `signal.NotifyContext` and passed to `exec.CommandContext`, so receiving a signal kills `say`
- Each iteration of the chunk loop checks `ctx.Err()` before proceeding to the next chunk

This makes long-running playback stoppable within ~1 second.

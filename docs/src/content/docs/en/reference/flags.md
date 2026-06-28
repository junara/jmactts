---
title: Flags
description: Every jmactts flag, exit code, and external dependency
---

## Flags

| short | long | Description |
|:--|:--|:--|
| `-v` | `--voice` | Voice name (e.g. `Kyoko` / `Samantha` / `Daniel`) |
| `-L` | `--lang` | Language/country code; auto-picks the primary voice when `-v` is absent |
| `-r` | `--rate` | Speech rate (words per minute) |
| `-f` | `--file` | Input text file (`-` for stdin) |
| `-c` | `--clipboard` | Read from `pbpaste` |
| `-o` | `--output` | Output audio file (format detected from extension) |
| `-l` | `--list-voices` | List available voices (filter with `-L`) |
| `-V` | `--version` | Show version |
| `-h` | `--help` | Show help |

Short and long forms are equivalent (`-v Kyoko` == `--voice Kyoko`).

## Exit codes

| Code | Meaning |
|--:|---|
| `0` | Success |
| `1` | Runtime error (voice not found / file read failure / `say` or `ffmpeg` error) |
| `2` | Flag parse error, or empty input |
| `130` | Interrupted with `Ctrl-C` (SIGINT) |

## External dependencies

| Command | Purpose | Required when |
|---|---|---|
| `/usr/bin/say` | Speech synthesis | Always (built into macOS) |
| `/usr/bin/pbpaste` | Clipboard read | `-c` is set (built into macOS) |
| `ffmpeg` | MP3 encoding | `-o *.mp3` is set |

## stdout / stderr

`-h` / `--help` and `-V` / `--version` go to **stdout**; usage messages on error go to **stderr**.

```bash
jmactts --help | less          # help goes to stdout, so piping works
jmactts --version              # stdout
```

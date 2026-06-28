---
title: Flags
description: Every jmactts flag in one place
---

## Reference

| short | long | Description |
|:--|:--|:--|
| `-v` | `--voice` | Voice name (e.g. `Kyoko` / `Samantha` / `Daniel`) |
| `-L` | `--lang` | Language/country code; auto-picks a voice when `-v` is absent |
| `-r` | `--rate` | Speech rate (words per minute) |
| `-f` | `--file` | Input text file (`-` for stdin) |
| `-c` | `--clipboard` | Read from `pbpaste` |
| `-o` | `--output` | Output audio file (format detected from extension) |
| `-l` | `--list-voices` | List available voices (filter with `-L`) |
| `-V` | `--version` | Show version |
| `-h` | `--help` | Show help |

## Exit codes

| Code | Meaning |
|--:|---|
| `0` | Success |
| `1` | Runtime error (voice not found, file read failure, `say`/`ffmpeg` error, …) |
| `2` | Flag parse error or empty input |
| `130` | Interrupted with `Ctrl-C` (`SIGINT`) |

## External dependencies

| Command | Purpose | Required |
|---|---|---|
| `/usr/bin/say` | Speech synthesis | Yes (built into macOS) |
| `/usr/bin/pbpaste` | Clipboard read | Only with `-c` |
| `ffmpeg` | MP3 encoding | Only when outputting MP3 |

## Help routing

`-h` / `--help` prints to **stdout**, while usage messages emitted on error go to **stderr**.

```bash
jmactts --help | less       # works because help goes to stdout
```

---
title: Input Sources
description: Read text from arguments, files, stdin, or the clipboard
---

`jmactts` supports four input paths. When several are present, precedence is:

```
-c (clipboard)  >  -f (file)  >  positional arguments  >  piped stdin
```

## Positional arguments

Reads the text passed as arguments. Multiple arguments are joined with spaces.

```bash
jmactts "Hello, world"
jmactts -L en "Use quotes for, punctuated, sentences."
```

## File (`-f`)

Use `-f` / `--file` for a text file.

```bash
jmactts -f speech.txt
jmactts -L en -f long_article.txt -o article.mp3
```

Pass `-f -` to read from stdin explicitly; this overrides positional arguments.

```bash
echo "Hello" | jmactts -f -
```

## Stdin (pipe)

When neither positional args, `-f`, nor `-c` are given and stdin is piped, jmactts reads stdin automatically.

```bash
echo "From stdin" | jmactts
cat article.txt | jmactts -L en
git log -1 --pretty=%s | jmactts -L en       # read the latest commit subject
```

## Clipboard (`-c`)

`-c` / `--clipboard` reads from the macOS clipboard via `pbpaste`.

```bash
jmactts -c
jmactts -c -L ja      # read in Japanese
```

Handy for piping an article you just copied from a browser into TTS.

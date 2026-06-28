---
title: Input Sources
description: Read text from arguments, files, stdin, or the clipboard
---

`jmactts` supports four input paths, evaluated in this order of precedence:

`-c` (clipboard) > `-f` (file) > positional arguments > piped stdin

## Positional arguments

Multiple arguments are joined with spaces.

```bash
jmactts "Hello, world"
```

## File

Use `-f` / `--file` to read from a text file.

```bash
jmactts -f speech.txt
```

`-f -` reads from stdin.

```bash
echo "Hello" | jmactts -f -
```

## Stdin (pipe)

When neither positional args nor `-f` are given and stdin is piped, jmactts reads stdin automatically.

```bash
echo "From stdin" | jmactts
cat article.txt | jmactts -L en
```

## Clipboard

`-c` / `--clipboard` reads from the macOS clipboard via `pbpaste`.

```bash
jmactts -c
jmactts -c -L ja
```

Handy when you just copied an article from your browser.

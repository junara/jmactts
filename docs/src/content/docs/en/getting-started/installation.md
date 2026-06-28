---
title: Installation
description: How to install jmactts
---

## Requirements

- macOS (uses `say`, `pbpaste`, `afconvert`)
- `ffmpeg` (only for MP3 output: `brew install ffmpeg`)

## Homebrew (recommended)

```bash
brew install junara/tap/jmactts
```

## Go install

```bash
go install github.com/junara/jmactts@latest
```

## Binary download

Download `darwin_amd64` / `darwin_arm64` tarballs from [Releases](https://github.com/junara/jmactts/releases).

## Build from source

```bash
git clone https://github.com/junara/jmactts.git
cd jmactts
go build -o jmactts .
sudo mv jmactts /usr/local/bin/
```

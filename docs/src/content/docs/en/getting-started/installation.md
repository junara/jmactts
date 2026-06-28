---
title: Installation
description: How to install jmactts
---

## Requirements

- macOS (uses `say` / `pbpaste` / `afconvert`)
- `ffmpeg` — only for MP3 output (`brew install ffmpeg`)

## Homebrew (recommended)

```bash
brew install junara/tap/jmactts
```

## Go install

Installs into `$GOPATH/bin`.

```bash
go install github.com/junara/jmactts@latest
```

## Binary download

Grab `darwin_amd64` (Intel Mac) or `darwin_arm64` (Apple Silicon) tarballs from [Releases](https://github.com/junara/jmactts/releases), extract them, and place the `jmactts` binary somewhere on your `PATH`.

## Build from source

```bash
git clone https://github.com/junara/jmactts.git
cd jmactts
go build -o jmactts .
sudo mv jmactts /usr/local/bin/
```

## Verify

```bash
jmactts --version
jmactts -L en "Hello, world"
```

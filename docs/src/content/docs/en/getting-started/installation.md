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

## Claude Code plugin (skill)

This repository doubles as a Claude Code **plugin marketplace** and ships a skill that documents how to use `jmactts`. Once installed, Claude can look up how to pick an input source, auto-select a voice with `-L`, export files with `-o`, and more, so it invokes `jmactts` correctly.

```text
/plugin marketplace add junara/jmactts
/plugin install jmactts@jmactts
```

The skill lives in [`plugins/jmactts/`](https://github.com/junara/jmactts/tree/main/plugins/jmactts), and `.claude-plugin/marketplace.json` is the catalog.

To update an installed skill, run the following (`/plugin marketplace update` only refreshes the catalog and does not update installed plugins):

```text
/plugin update jmactts@jmactts
```

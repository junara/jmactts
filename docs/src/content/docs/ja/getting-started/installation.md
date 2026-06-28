---
title: インストール
description: jmactts のインストール手順
---

## 必要条件

- macOS (`say` / `pbpaste` / `afconvert` を利用)
- `ffmpeg` — MP3 出力時のみ (`brew install ffmpeg`)

## Homebrew (推奨)

```bash
brew install junara/tap/jmactts
```

## Go install

`$GOPATH/bin` にインストールされます。

```bash
go install github.com/junara/jmactts@latest
```

## バイナリダウンロード

[Releases](https://github.com/junara/jmactts/releases) から `darwin_amd64` (Intel Mac) または `darwin_arm64` (Apple Silicon) の tar.gz を取得し、展開後に `jmactts` バイナリを PATH に配置します。

## ソースからビルド

```bash
git clone https://github.com/junara/jmactts.git
cd jmactts
go build -o jmactts .
sudo mv jmactts /usr/local/bin/
```

## 動作確認

```bash
jmactts --version
jmactts -L ja こんにちは
```

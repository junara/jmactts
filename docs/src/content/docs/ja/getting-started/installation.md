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

## Claude Code プラグイン (skill)

このリポジトリは Claude Code の**プラグインマーケットプレイス**を兼ねており、`jmactts` の使い方をまとめた skill を配布しています。取り込むと、Claude が入力ソースの選び方や `-L` によるボイス自動選択、`-o` でのファイル出力などを参照して `jmactts` を正しく呼び出せるようになります。

```text
/plugin marketplace add junara/jmactts
/plugin install jmactts@jmactts
```

skill は [`plugins/jmactts/`](https://github.com/junara/jmactts/tree/main/plugins/jmactts) にあり、`.claude-plugin/marketplace.json` がカタログです。

skill を更新するには次を実行します (`/plugin marketplace update` はカタログを更新するだけで、インストール済みの skill は更新されません)。

```text
/plugin update jmactts@jmactts
```

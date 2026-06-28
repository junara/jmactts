# jmactts

[ドキュメントサイト](https://junara.github.io/jmactts/) | [English](https://junara.github.io/jmactts/en/)

macOS 標準の `say` コマンドをラップした、多言語対応のテキスト読み上げ CLI。

- 引数 / ファイル / 標準入力 / クリップボードから読み上げ
- 国・言語コードからのプライマリボイス自動選択 (`ja` / `JP` / `ja_JP`)
- AIFF / WAV / M4A / AAC / MP3 のファイル出力 (MP3 は `ffmpeg` 経由)
- 長文の自動分割再生と Ctrl-C による即時停止

## 要件

- macOS (`say` / `pbpaste` / `afconvert` を利用)
- `ffmpeg` — MP3 出力時のみ (`brew install ffmpeg`)
- Go 1.26+ — ソースからビルドする場合のみ

## インストール

```sh
# Homebrew (推奨)
brew install junara/tap/jmactts

# Go
go install github.com/junara/jmactts@latest

# バイナリ
# https://github.com/junara/jmactts/releases から darwin_amd64 / darwin_arm64 を取得

# ソース
go build -o jmactts . && sudo mv jmactts /usr/local/bin/
```

## クイックスタート

```sh
jmactts こんにちは 世界                          # 引数を読み上げ
echo "Hello" | jmactts -v Samantha               # 標準入力 + ボイス指定
jmactts -c                                        # クリップボードを読み上げ
jmactts -L ja こんにちは                          # 言語コードでボイス自動選択
jmactts -L ja -o hello.mp3 ここに本文            # MP3 で保存 (要 ffmpeg)
jmactts -l -L ja                                  # 日本語ボイス一覧
```

より詳しい使い方は [ドキュメントサイト](https://junara.github.io/jmactts/) を参照してください。

## フラグ

| short | long | 説明 |
|:--|:--|:--|
| `-v` | `--voice` | ボイス名 (例: `Kyoko` / `Samantha` / `Daniel`) |
| `-L` | `--lang` | 国・言語コード。`-v` 未指定時にプライマリボイスを自動選択 |
| `-r` | `--rate` | 話速 (words per minute) |
| `-f` | `--file` | 入力テキストファイル (`-` で標準入力) |
| `-c` | `--clipboard` | `pbpaste` から入力 |
| `-o` | `--output` | 音声ファイル出力 (拡張子で形式を自動判定) |
| `-l` | `--list-voices` | 利用可能なボイス一覧 (`-L` で絞り込み可) |
| `-V` | `--version` | バージョン表示 |
| `-h` | `--help` | ヘルプ |

入力ソースの優先順位は `-c` > `-f` > 位置引数 > パイプされた標準入力。

## アーキテクチャ

Clean Architecture の 3 層構成 (`domain` / `usecase` / `adapter`)。依存方向は内向き一方向で、`usecase` がポート (インターフェース) を所有します。

新しい入力ソースや出力フォーマットを追加したい場合は、対応するアダプタを実装するだけで `usecase` / `domain` を変更する必要はありません。

詳細は [ドキュメントサイトのアーキテクチャ](https://junara.github.io/jmactts/ja/reference/architecture/) を参照してください。

## 開発

```sh
go test ./...
go build -o jmactts .
```

## リリース

`v*` タグの push で [`.github/workflows/release.yml`](.github/workflows/release.yml) が起動し、[goreleaser](https://goreleaser.com/) が以下を実行します。

- darwin (amd64 / arm64) のバイナリビルドとアーカイブ作成
- GitHub Releases へアセットとチェックサムを公開
- `junara/homebrew-tap` リポジトリの **Homebrew Cask** を自動更新

リリース前に Secret `HOMEBREW_TAP_GITHUB_TOKEN` (tap リポジトリへの write 権限を持つ PAT) を本リポジトリに登録しておきます。

```sh
git tag v0.1.0
git push origin v0.1.0
```

## ライセンス

[MIT License](LICENSE) — Copyright (c) 2026 junara

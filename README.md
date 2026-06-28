# jmactts

macOS 標準の `say` コマンドをラップした、多言語対応のテキスト読み上げ CLI。

- 引数 / ファイル / 標準入力 / クリップボードから読み上げ
- 国/言語コードで自動ボイス選択 (`ja` / `JP` / `ja_JP` など)
- AIFF / WAV / M4A / AAC / MP3 出力 (MP3 は `ffmpeg` 経由)
- 長文の自動分割再生、Ctrl-C で即停止

## 要件

- macOS (`say`, `pbpaste`, `afconvert` を利用)
- Go 1.26+ (ビルド時のみ)
- `ffmpeg` (MP3 出力時のみ。`brew install ffmpeg`)

## インストール

### Homebrew (推奨)

```sh
brew install junara/tap/jmactts
```

### Go install

```sh
go install github.com/junara/jmactts@latest
```

### バイナリダウンロード

[Releases](https://github.com/junara/jmactts/releases) から `darwin_amd64` / `darwin_arm64` の tar.gz を取得。

### ソースからビルド

```sh
go build -o jmactts .
sudo mv jmactts /usr/local/bin/
```

## 使い方

```sh
# 引数で渡す
jmactts こんにちは 世界

# 標準入力 / ファイル
echo "Hello, world" | jmactts -v Samantha
jmactts -f speech.txt

# クリップボードの内容を読み上げ
jmactts -c

# 国/言語コードで自動ボイス選択
jmactts -L ja_JP こんにちは            # 完全ロケール
jmactts -L ja こんにちは                # 言語コード
jmactts -L JP こんにちは                # 国コード
jmactts -L en Hello world              # en_* のいずれか

# 話速調整
jmactts -v Kyoko -r 200 ゆっくり読ませたいテキスト

# 音声ファイル出力 (拡張子で形式を自動判定)
jmactts -o hello.aiff こんにちは        # AIFF (say の既定)
jmactts -o hello.wav  こんにちは        # WAVE / LEI16 22050Hz
jmactts -o hello.m4a  こんにちは        # AAC in M4A
jmactts -o hello.mp3  こんにちは        # libmp3lame (要 ffmpeg)

# ボイス一覧
jmactts -l                              # 全ボイス (sample 付き)
jmactts -l -L ja                        # 日本語のみ
```

## フラグ

| short | long | 説明 |
|:--|:--|:--|
| `-v` | `--voice` | ボイス名 (例: `Kyoko` / `Samantha` / `Daniel`) |
| `-L` | `--lang` | 国/言語コード。`-v` 未指定時に自動選択 |
| `-r` | `--rate` | 話速 (words per minute) |
| `-f` | `--file` | 入力テキストファイル (`-` で標準入力) |
| `-c` | `--clipboard` | `pbpaste` から入力 |
| `-o` | `--output` | 音声ファイル出力 (拡張子で形式判定) |
| `-l` | `--list-voices` | 利用可能なボイス一覧 (`-L` で絞り込み可) |
| `-V` | `--version` | バージョン表示 |
| `-h` | `--help` | ヘルプ |

入力ソースの優先順位: `-c` > `-f` > 位置引数 > パイプされた標準入力。

## 国/言語コードの照合

`-L` の値は以下の順で照合します。

1. `ja_JP` 形式 (アンダースコア含む) → ロケール完全一致
2. 言語コード (`ja`, `en` 等) → `xx_*` の全マッチ
3. それでもマッチしない場合は国コード (`JP`, `US` 等) として `*_YY` を探索

複数マッチした場合、**ボイス名にカッコ `(` を含まないもの**を優先します (macOS の慣習で、各言語のプライマリボイスは `Kyoko`/`Samantha`/`Daniel` 等、カッコなしの命名のため)。

## 長文 / 中断

1500 文字を超える入力は、句点 (`。．.!?！？` / 改行) で分割して順次 `say` に渡します (ファイル出力時は分割しません)。

再生中の Ctrl-C (`SIGINT`) で即座に停止し、終了コード 130 を返します。

## アーキテクチャ

Clean Architecture の 3 層構成。依存方向は内向き一方向で、`usecase` がポート (インターフェース) を所有します。

```
main.go                              ← Composition Root
└─ internal/
   ├─ domain/                        ← Entities (純粋ドメイン, 外部依存なし)
   │  ├─ voice.go                    Voice / Locale / VoiceList
   │  ├─ speech.go                   OutputFormat / DetectFormat
   │  └─ chunk.go                    ChunkText (純粋関数)
   ├─ usecase/                       ← Application Logic
   │  ├─ ports.go                    Synthesizer / MP3Encoder / VoiceCatalog / TextSource
   │  ├─ speak.go                    SpeakUseCase
   │  └─ voices.go                   ListVoices / PickVoice
   └─ adapter/                       ← Interface Adapters
      ├─ say/                        macOS say を Synthesizer + VoiceCatalog として実装
      ├─ ffmpeg/                     ffmpeg を MP3Encoder として実装
      ├─ textsource/                 Clipboard / File / Stdin / Args (TextSource 実装)
      └─ cli/                        フラグ解析、Run(ctx, args, deps, stdout, stderr) int
```

新しい入力ソースや出力形式は、対応するアダプタを追加するだけで `usecase` / `domain` を変更せずに拡張できます。

## 開発

```sh
go test ./...
go build -o jmactts .
```

## リリース

`v*` タグの push で [`.github/workflows/release.yml`](.github/workflows/release.yml) が起動し、
[goreleaser](https://goreleaser.com/) が以下を実行します。

- darwin (amd64 / arm64) のバイナリビルド + アーカイブ
- GitHub Releases へアセットとチェックサムを公開
- `junara/homebrew-tap` リポジトリへ Formula を自動 PR / push (Homebrew tap 更新)

リリース前に、本リポジトリ側に Secret `HOMEBREW_TAP_GITHUB_TOKEN` (tap リポジトリへの write 権限を持つ PAT) を設定してください。

```sh
git tag v0.1.0
git push origin v0.1.0
```

## ライセンス

[MIT License](LICENSE) — Copyright (c) 2026 junara

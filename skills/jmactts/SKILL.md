---
name: jmactts
description: macOS の `say` コマンドをラップした多言語テキスト読み上げ CLI `jmactts` を使うためのガイド。引数 / ファイル / 標準入力 / クリップボード (`pbpaste`) からの読み上げ、`-L` による国・言語コードからの自動ボイス選択、AIFF / WAV / M4A / AAC / MP3 へのファイル出力、長文の自動分割再生、Ctrl-C による即時停止などをカバー。「テキストを読み上げて」「日本語を音声化」「文章を MP3 にして」のような依頼に使う。
---

# jmactts — macOS テキスト読み上げ skill

`/usr/bin/say` をラップした多言語対応 CLI。引数・ファイル・標準入力・クリップボードからのテキストを再生したり音声ファイルに保存したりできる。

## 前提条件

- macOS (`say`, `pbpaste`, `afconvert` を利用)
- `jmactts` がインストール済みであること
  - `brew install junara/tap/jmactts` (推奨)
  - もしくは `go install github.com/junara/jmactts@latest`
- MP3 出力時のみ `ffmpeg` が必要 (`brew install ffmpeg`)

未インストールなら最初にユーザーへ案内する。

## 基本構文

```
jmactts [flags] [text...]
```

- 位置引数があればそのテキストを読み上げる
- 引数がなく `-f` も `-c` も無く stdin がパイプされていれば stdin を読む
- `-v` でボイス名、`-L` で言語/国コード、`-o` でファイル保存
- `-V` でバージョン、`-h` でヘルプ

## 入力ソース

優先順位は **`-c` > `-f` > 位置引数 > パイプ stdin**。

```bash
# 位置引数
jmactts こんにちは 世界

# ファイル
jmactts -f speech.txt
jmactts -f - <<< "標準入力からのテスト"   # -f - で stdin

# パイプ stdin
echo "標準入力です" | jmactts
cat article.txt | jmactts -L ja

# クリップボード (pbpaste)
jmactts -c
jmactts -c -L ja                          # クリップボード内容を日本語ボイスで
```

「クリップボードの内容を読み上げて」と言われたら `jmactts -c` を使う。

## ボイス選択

### 言語/国コードで自動選択 (`-L`) ← 第一推奨

`-v` を指定せず `-L` だけ与えると、その言語/国のプライマリボイスが自動で選ばれる。**ユーザーがボイス名を明示しない場合は常にこちらを使う**。

```bash
jmactts -L ja_JP こんにちは      # 完全ロケール → Kyoko
jmactts -L ja こんにちは          # 言語コードのみ → Kyoko
jmactts -L JP こんにちは          # 国コードのみ → Kyoko
jmactts -L en Hello world         # 英語のプライマリ → Samantha
jmactts -L en_GB Hello, mate      # イギリス英語 → Daniel
```

照合順:

1. `ja_JP` 形式 (アンダースコア含む) → ロケール完全一致
2. 言語コード (`ja`, `en`, …) → `xx_*` の全マッチ
3. 国コード (`JP`, `US`, …) → `*_YY` の全マッチ

複数マッチした場合は **ボイス名にカッコ `(` を含まないもの (=プライマリ)** が優先される。

### ボイス名で指定 (`-v`)

`say -v ?` の一覧から好みのボイスを直接指定する。`-v` は `-L` より優先される。

```bash
jmactts -v Kyoko こんにちは
jmactts -v Samantha "Hello there"
jmactts -v Otoya こんばんは
```

### ボイス一覧

```bash
jmactts -l                       # 全ボイス (sample 付き)
jmactts -l -L ja                 # 日本語ボイスのみ絞り込み
jmactts -l -L en_GB              # 英国英語のみ
```

### 話速

```bash
jmactts -L ja -r 250 早口で読み上げる
jmactts -L ja -r 120 ゆっくり読み上げる
```

## 出力フォーマット

`-o` の **拡張子から自動判定** される。

| 拡張子 | フォーマット | 経路 |
|---|---|---|
| `.aiff` / `.aif` | AIFF (PCM, `say` の既定) | `say` 単体 |
| `.wav` | WAVE / LEI16 22050Hz mono | `say` 単体 |
| `.m4a` / `.aac` | AAC in M4A | `say` 単体 |
| `.mp3` | MPEG Layer III (libmp3lame VBR Q2) | 一時 AIFF → `ffmpeg` |

```bash
jmactts -L ja -o hello.aiff こんにちは
jmactts -L ja -o hello.wav  こんにちは
jmactts -L ja -o hello.m4a  こんにちは
jmactts -L ja -o hello.mp3  こんにちは           # 要 ffmpeg
```

「読み上げたものを mp3 で保存して」と言われたら `-o ファイル名.mp3` を付ける。

## 長文と中断

- **1500 文字超は文区切り (`。．.!?！？` / 改行) で自動分割**して順次再生 (再生時のみ。ファイル出力時は分割しない)
- 再生中の **`Ctrl-C` (SIGINT) で即座に停止**、終了コードは 130

長文ファイルを読ませる場合に追加設定は不要。

```bash
cat long_article.txt | jmactts -L ja          # 自動分割
jmactts -f long_article.txt -L ja -r 180      # ファイル + 速度指定
```

## フラグ一覧

| short | long | 説明 |
|:--|:--|:--|
| `-v` | `--voice` | ボイス名 (例: `Kyoko` / `Samantha` / `Daniel`) |
| `-L` | `--lang` | 国/言語コード。`-v` 未指定時に自動選択 |
| `-r` | `--rate` | 話速 (words per minute) |
| `-f` | `--file` | 入力テキストファイル (`-` で stdin) |
| `-c` | `--clipboard` | `pbpaste` から入力 |
| `-o` | `--output` | 音声ファイル出力 (拡張子で形式判定) |
| `-l` | `--list-voices` | 利用可能なボイス一覧 (`-L` で絞り込み可) |
| `-V` | `--version` | バージョン表示 |
| `-h` | `--help` | ヘルプ |

### 終了コード

| コード | 意味 |
|--:|---|
| `0` | 正常終了 |
| `1` | 実行時エラー (ボイス未検出、`say` / `ffmpeg` のエラー等) |
| `2` | フラグ解析エラー、または読み上げ対象が空 |
| `130` | `Ctrl-C` で中断 |

## よくある用途

### 1. 渡された文章をそのまま読ませる

ユーザーが「これを読み上げて」「音声化して」と言ってきたら:

```bash
jmactts -L ja "渡された文章をそのまま"
```

長い場合や改行を含む場合はヒアドキュメント:

```bash
jmactts -L ja <<'EOF'
複数行の
文章でも
そのまま読めます。
EOF
```

### 2. ブラウザでコピーした記事を読む

```bash
jmactts -c -L ja
```

### 3. 音声ファイルを生成する (mp3)

```bash
jmactts -L ja -o output.mp3 ここに本文を書く
# あるいはファイルから
jmactts -L ja -f article.txt -o article.mp3
```

### 4. 英語テキストをイギリス英語で

```bash
jmactts -L en_GB "Good evening, my dear chap."
```

### 5. ゆっくり再生して聞き取り練習

```bash
jmactts -L en -r 130 -f english_passage.txt
```

### 6. 利用可能なボイスを確認

```bash
jmactts -l -L ja
jmactts -l | head -50
```

### 7. パイプラインの一部として

```bash
# 直近コミットのタイトルを読み上げ
git log -1 --pretty=%s | jmactts -L ja

# 何らかの処理結果をクリップボード経由で
some-command | pbcopy && jmactts -c
```

## トラブルシューティング

- **「ボイスが見つかりません」**: `jmactts -l` で利用可能なボイス名を確認し、`-v` で正しい名前を指定する。`-L` 経由なら typo してもプライマリボイスが選ばれる
- **MP3 出力で失敗**: `which ffmpeg` を確認。無ければ `brew install ffmpeg`
- **クリップボードが空**: `pbpaste` が何も返していない。先に `pbcopy` するか、内容を直接渡す
- **長文で再生が遅い / 止まらない**: チャンク境界で待つ必要あり。`Ctrl-C` で即停止する。それでも問題ならファイル出力 (`-o`) に切り替える (分割しない)

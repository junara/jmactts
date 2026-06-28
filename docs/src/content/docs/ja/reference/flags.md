---
title: フラグ一覧
description: jmactts のすべてのフラグ・終了コード・外部依存
---

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

短形と長形は等価です (例: `-v Kyoko` と `--voice Kyoko` は同じ)。

## 終了コード

| コード | 意味 |
|--:|---|
| `0` | 正常終了 |
| `1` | 実行時エラー (ボイス未検出 / ファイル読み込み失敗 / `say` や `ffmpeg` のエラー等) |
| `2` | フラグ解析エラー、または読み上げ対象が空 |
| `130` | `Ctrl-C` (SIGINT) で中断 |

## 外部依存

| コマンド | 用途 | 必要となる条件 |
|---|---|---|
| `/usr/bin/say` | 音声合成 | 常時必須 (macOS 標準) |
| `/usr/bin/pbpaste` | クリップボード読み取り | `-c` 指定時のみ (macOS 標準) |
| `ffmpeg` | MP3 エンコード | `-o *.mp3` 指定時のみ |

## stdout / stderr の使い分け

`-h` / `--help` と `-V` / `--version` の出力は **stdout** に、エラー時の使用法表示は **stderr** に出力されます。

```bash
jmactts --help | less          # stdout に出るので less に渡せる
jmactts --version              # stdout
```

---
title: 出力フォーマット
description: AIFF / WAV / M4A / AAC / MP3 でファイルに書き出す
---

`-o` / `--output` でファイル出力できます。フォーマットは **拡張子から自動判定** されます。

| 拡張子 | フォーマット | 経路 |
|---|---|---|
| `.aiff` / `.aif` | AIFF (PCM、`say` の既定) | `say` 単体 |
| `.wav` | WAVE / LEI16 22050 Hz mono | `say` 単体 |
| `.m4a` / `.aac` | AAC in M4A | `say` 単体 |
| `.mp3` | MPEG Layer III (libmp3lame VBR Q2) | 一時 AIFF → `ffmpeg` |

## 例

```bash
jmactts -L ja -o hello.aiff こんにちは
jmactts -L ja -o hello.wav  こんにちは
jmactts -L ja -o hello.m4a  こんにちは
jmactts -L ja -o hello.mp3  こんにちは        # 要 ffmpeg
```

## MP3 出力の仕組み

`say` は MP3 を直接出力できないため、`jmactts` は内部で次の処理を行います。

1. `say` で一時 AIFF ファイルを生成
2. `ffmpeg -codec:a libmp3lame -qscale:a 2` で MP3 に変換
3. 一時ファイルを削除

そのため MP3 出力には **`ffmpeg` のインストールが必須** です。

```bash
brew install ffmpeg
```

未インストール時は次のエラーで終了します。

```
jmactts: MP3 出力には ffmpeg が必要です: brew install ffmpeg
```

## ファイル出力時の補足

- 長文を渡してもファイル出力時は **分割せず単一ファイル** を生成します ([長文と中断](/jmactts/ja/usage/long-text/) 参照)
- 既存ファイルは上書きされます

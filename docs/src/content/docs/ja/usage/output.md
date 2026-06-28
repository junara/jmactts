---
title: 出力フォーマット
description: AIFF / WAV / M4A / AAC / MP3 のファイル出力
---

`-o` / `--output` でファイル出力できます。フォーマットは **拡張子から自動判定** されます。

| 拡張子 | フォーマット | 経路 |
|---|---|---|
| `.aiff` / `.aif` | AIFF (PCM、`say` の既定) | `say -o` |
| `.wav` | WAVE / LEI16 22050 Hz mono | `say --file-format=WAVE` |
| `.m4a` / `.aac` | AAC in M4A | `say --file-format=m4af --data-format=aac` |
| `.mp3` | MPEG Audio Layer III (libmp3lame VBR Q2) | AIFF 一時生成 → `ffmpeg` |

## 例

```bash
jmactts -L ja -o hello.aiff こんにちは
jmactts -L ja -o hello.wav  こんにちは
jmactts -L ja -o hello.m4a  こんにちは
jmactts -L ja -o hello.mp3  こんにちは
```

## MP3 出力について

MP3 は `say` 単体では生成できないため、`jmactts` は内部で次の処理を行います。

1. `say --data-format=...` で一時 AIFF ファイルを生成
2. `ffmpeg -codec:a libmp3lame -qscale:a 2` で MP3 に変換
3. 一時ファイルを削除

そのため MP3 出力には **`ffmpeg` のインストールが必須** です。

```bash
brew install ffmpeg
```

不在時は以下のエラーで停止します。

```
jmactts: MP3 出力には ffmpeg が必要です: brew install ffmpeg
```

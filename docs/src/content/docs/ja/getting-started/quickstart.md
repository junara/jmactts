---
title: クイックスタート
description: jmactts を 5 分で使い始める
---

インストール後、すぐに使い始められます。

## まず読み上げてみる

```bash
jmactts こんにちは 世界
```

## 言語を指定して読み上げ

`-L` で言語/国コードを指定すると、その言語のプライマリボイスが自動選択されます。

```bash
jmactts -L ja Hello は日本語ボイスで読まれます
jmactts -L en Hello will be read in English
```

## ファイル出力

拡張子から音声フォーマットを自動判定します。

```bash
jmactts -L ja -o hello.m4a こんにちは
jmactts -L ja -o hello.mp3 MP3 でも保存できます
```

## クリップボード読み上げ

`pbcopy` でコピーしたテキストを即座に再生:

```bash
pbpaste | jmactts        # パイプ経由
jmactts -c               # フラグでも同じ
```

## 次のステップ

- [入力ソース](/jmactts/ja/usage/input/) — 引数 / ファイル / 標準入力 / クリップボード
- [ボイス選択](/jmactts/ja/usage/voice/) — `-v` と `-L` の使い分け
- [出力フォーマット](/jmactts/ja/usage/output/) — AIFF / WAV / M4A / MP3
- [フラグ一覧](/jmactts/ja/reference/flags/) — すべてのフラグ

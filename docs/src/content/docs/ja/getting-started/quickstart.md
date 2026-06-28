---
title: クイックスタート
description: jmactts を 5 分で使い始める
---

インストール後、すぐに使い始められます。

## まず読み上げてみる

```bash
jmactts こんにちは 世界
```

`say` の既定ボイス (システム設定で選んだもの) で読み上げられます。

## 言語を指定して読み上げ

`-L` に言語または国コードを渡すと、その言語のプライマリボイスが自動で選ばれます。

```bash
jmactts -L ja こんにちは            # Kyoko で読み上げ
jmactts -L en "Hello, world"        # Samantha で読み上げ
jmactts -L en_GB "Good evening"     # Daniel で読み上げ (イギリス英語)
```

## ファイル出力

`-o` の拡張子から音声フォーマットが自動判定されます。

```bash
jmactts -L ja -o hello.m4a こんにちは
jmactts -L ja -o hello.mp3 MP3 でも保存できます   # 要 ffmpeg
```

## クリップボードを読み上げ

ブラウザでコピーした記事をそのまま音声化:

```bash
jmactts -c             # フラグで指定
pbpaste | jmactts      # パイプ経由でも同じ
```

## 次のステップ

- [入力ソース](/jmactts/ja/usage/input/) — 引数 / ファイル / 標準入力 / クリップボードの使い分け
- [ボイス選択](/jmactts/ja/usage/voice/) — `-v` と `-L` の使い分けと照合ロジック
- [出力フォーマット](/jmactts/ja/usage/output/) — AIFF / WAV / M4A / AAC / MP3 の使い分け
- [フラグ一覧](/jmactts/ja/reference/flags/) — すべてのフラグと終了コード

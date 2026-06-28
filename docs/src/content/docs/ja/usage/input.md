---
title: 入力ソース
description: 引数 / ファイル / 標準入力 / クリップボードからテキストを読み込む
---

`jmactts` は 4 つの入力経路をサポートします。優先順位は以下のとおりです。

`-c` (クリップボード) > `-f` (ファイル) > 位置引数 > パイプされた標準入力

## 位置引数

複数の引数はスペース連結されて読み上げられます。

```bash
jmactts こんにちは 世界
```

## ファイル

`-f` / `--file` でテキストファイルを指定します。

```bash
jmactts -f speech.txt
```

`-f -` を指定すると標準入力から読み込みます。

```bash
echo "Hello" | jmactts -f -
```

## 標準入力 (パイプ)

引数も `-f` も無く、stdin がパイプされている場合は自動的に標準入力を読みます。

```bash
echo "標準入力からのテスト" | jmactts
cat article.txt | jmactts -L ja
```

## クリップボード

`-c` / `--clipboard` で macOS のクリップボード (`pbpaste`) から読み込みます。

```bash
jmactts -c
jmactts -c -L ja
```

ブラウザでコピーした記事をそのまま読み上げたいときに便利です。
